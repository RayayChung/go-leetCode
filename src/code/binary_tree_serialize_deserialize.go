package main

import (
	"fmt"
	"strings"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

//func Constructor() Codec {
//	return Codec{}
//}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var queue []*TreeNode
	var result []string
	queue = append(queue, root)

	var emptyNode *TreeNode
	for len(queue) != 0 {
		node := queue[0]
		fmt.Println(node)
		if node != nil {
			result = append(result, string(node.Val))
		} else {
			result = append(result, "null")
			continue
		}

		if node.Left == nil {
			queue = append(queue, emptyNode)
		} else {
			queue = append(queue, node.Left)
		}

		if node.Right == nil {
			queue = append(queue, emptyNode)
		} else {
			queue = append(queue, node.Right)
		}

		queue = queue[1:]
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//
//}

//func main() {
//	obj := Constructor()
//
//	root := &TreeNode{
//		1,
//		&TreeNode{
//			2,
//			nil,
//			nil,
//		},
//		&TreeNode{
//			3,
//			nil,
//			nil,
//		},
//	}
//
//	str := obj.serialize(root)
//	fmt.Print(str)
//}
