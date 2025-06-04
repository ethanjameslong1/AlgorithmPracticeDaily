package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		temp.Next = &ListNode{}
		temp = temp.Next
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
		temp.Val = val2 + val1 + carry
		carry = temp.Val / 10
		temp.Val = temp.Val % 10
	}

	return dummyHead.Next
}
