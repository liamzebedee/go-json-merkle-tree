package main

import (
	"fmt"
	"testing"
	// "fmt"
	"github.com/stretchr/testify/assert"
)

// func TestMakesTree(t *testing.T) {
// 	loadFile("./test.json")
// }

var sampleMap = make(map[string]interface{})

func TestMain(m *testing.M) { 
	sampleMap["0"] = "Hi"
	m.Run()
}

// func TestToTreeMap(t *testing.T) {
	
// 	tree := toTree(sample)

	// expect := node{
	// 	children: []node{
	// 		node{
	// 			val: getBytes("0"),
	// 			children: []node{
	// 				node{
	// 					val: getBytes("Hi"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }

// 	assert.Equal(t, expect, tree)
// }

func TestToTreeArray(t *testing.T) {
	var sample []interface{}
	sample = append(sample, sampleMap)
	sample = append(sample, 12)

	tree := toTree(sample)

	expect := node{
		children: []node{
			node{
				children: []node{
					node{
						val: getBytes("0"),
						children: []node{
							node{
								val: getBytes("Hi"),
							},
						},
					},
				},
			},

			node{
				val: getBytes(12),
			},
		},
	}

	assert.Equal(t, expect, tree)
}

func TestCompareTree(t *testing.T) {

}
/*

	// tree.traverse(func(x node) {
	// 	fmt.Println(string(x.val))
	// })

*/