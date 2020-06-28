package main

/**
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	current := head.Next
	existMap := make(map[int]bool)

	existMap[pre.Val] = true
	for current != nil {
		key := current.Val
		if !existMap[key] {
			existMap[key] = true
			pre = current
			current = current.Next
		} else {
			pre.Next = current.Next
			current = current.Next
		}
	}
	return head
}
