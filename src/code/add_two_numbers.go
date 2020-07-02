package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{
		0,
		nil,
	}
	moveHead := head
	flag := 0
	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + flag

		var val int
		if sum >= 10 {
			flag = 1
			val = sum - 10
		} else {
			flag = 0
			val = sum
		}

		tmp := &ListNode{
			val,
			nil,
		}
		moveHead.Next = tmp
		moveHead = tmp
	}
	return head.Next
}

//func main() {
//	l1 := &ListNode{
//		1,
//		&ListNode{
//			8,
//			nil,
//		},
//	}
//
//	l2 := &ListNode{
//		0,
//		nil,
//	}
//
//	node := addTwoNumbers(l1, l2)
//
//	for node != nil {
//		fmt.Print(node.Val)
//		node = node.Next
//	}
//}
