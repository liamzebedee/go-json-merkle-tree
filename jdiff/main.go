package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"bytes"
	"os"
	"io/ioutil"
	"encoding/gob"
	"crypto/sha1"
	// "encoding/base64"
	"reflect"
)

var hasher = sha1.New()

type node struct {
	val []byte
	children []node
}

func (n node) traverse(visitor func(x node, depth int) bool, depth int) {
	if visitor(n, depth) == true {
		for _, child := range n.children {
			child.traverse(visitor, depth+1)
		}
	}
}

func (n node) hash() []byte {
	var x []byte
	x = append(x, n.val...)
	if len(n.children) > 0 {
		for _, y := range n.children {
			x = append(x, y.hash()...)
		}
	}
	return hasher.Sum(x)
}

func main() {
	f1 := os.Args[1]
	f2 := os.Args[2]
	
	fmt.Println(compareTree(loadFile(f1), loadFile(f2), 0))

	// x.traverse(func(x node, depth int) bool {
	// 	// x.hash = hasher(append(x.val, ...x.children.hash))
	// 	str := base64.StdEncoding.EncodeToString(x.hash())
	// 	indent := ""
	// 	for i := 0; i < depth; i++ {
	// 		indent += "\t"
	// 	}
	// 	fmt.Println(indent, str)
	// 	return true
	// 	// fmt.Println()
	// 	// fmt.Println(x.hash())
	// }, 0)

	// ioutil.ReadFile(f2)
	// loadFile(f1)
}

func loadFile(path string) node {
	if data, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		return processFile(data)
	}
}

func processFile(data []byte) node {
	var payload interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(data, &payload)

	x := toTree(payload)

	return x
}


func getBytes(key interface{}) ([]byte) {
	if key == nil {
		return []byte{0}
	}
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        panic(err)
    }
    return buf.Bytes()
}

const DEFAULT_SIZE = 16

func toTree(payload interface{}) node {
	if x, ok := payload.(map[string]interface{}); ok {
		n := node{
			val: nil,
			children: []node{},
		}

		for k, v := range x {
			n2 := node{
				val: getBytes(k),
				children: []node{
					toTree(v),
				},
			}
			n.children = append(n.children, n2)
		}
		return n
	}

	switch reflect.ValueOf(payload).Kind() {
	case reflect.Slice:
		n := node{
			children: []node{},
		}

		for _, y := range payload.([]interface{}) {
			child := toTree(y)
			n.children = append(n.children, child)
		}

		return n
	default:
		n := node{
			val: getBytes(payload),
		}
		return n
	}
	return node{}
}

func compare(n1, n2 node) bool {
	return bytes.Equal(n1.hash(), n2.hash())
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


func compareTree(n1, n2 node, d int) []node {
	var diffs []node

	if(!compare(n1, n2)) {
		diffs = append(diffs, n1)
		return diffs
	}

	for i := 0; i < min(len(n1.children), len(n2.children)); i++ {
		diffs = append(diffs, compareTree(n1.children[i], n2.children[i], d+1) ...)
	}

	return diffs
}

// func isEqual(n1, n2 node) bool {
// 	retval := true

// 	if n1.hash() != n2.hash() {
// 		fmt.Println()
// 		return false
// 	}
	// for i, _ := range math.Max(len(n1.children), len(n2.children)) {
	// 	if !isEqual(n1.children[i], n2.children[i]) {
	// 		retval = false
	// 	}
	// }

// 	return retval
// }