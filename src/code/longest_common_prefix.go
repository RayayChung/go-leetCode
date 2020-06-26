package main

import "fmt"

type TriTreeNode struct {
	val      string
	Children *TriTreeNode
}

func NewTreeNode(val string) *TriTreeNode {
	return &TriTreeNode{
		val,
		nil,
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	size := -1

	for _, val := range strs {
		if len(val) == 0 {
			return ""
		}
		if size == -1 {
			size = len(val)
		} else {
			size = min(size, len(val))
		}
	}

	first := strs[0]

	trie := NewTreeNode(first[0:1])
	root := trie
	for i := 1; i < size; i++ {
		childNode := NewTreeNode(first[i : i+1])
		trie.Children = childNode
		trie = childNode
	}

	same := size
	for i := 1; i < len(strs); i++ {

		str := strs[i]
		fmt.Printf("str~~~: %s\n", str)
		tmp := root
		subSame := 0

		for j := 0; j < size; j++ {
			if tmp.val == string(str[j]) {
				tmp = tmp.Children
				subSame++
				continue
			}
			break
		}
		if subSame == 0 {
			return ""
		}
		if same == -1 {
			same = subSame
		} else {
			same = min(same, subSame)
		}
	}

	var result string
	for i := 0; i < same; i++ {
		result += root.val
		root = root.Children
	}
	return result
}
