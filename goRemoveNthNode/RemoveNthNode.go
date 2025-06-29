package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	removeNthFromEnd(h, 2)

	for h.Next != nil {
		fmt.Println(h.Val)
		h = h.Next
		if h.Next == nil {
			fmt.Println(h.Val)
		}
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cNode := head
	count := 0
	if head.Next == nil {
		return nil
	}
	for cNode.Next != nil {
		count++
		cNode = cNode.Next
		if cNode.Next == nil {
			count++
		}
	}
	if count == n {
		return head.Next
	}
	count -= n - 1
	cNode = head
	for range count - 2 {
		cNode = cNode.Next
	}
	if cNode.Next.Next == nil {
		cNode.Next = nil
	} else {
		cNode.Next = cNode.Next.Next
	}

	return head
}
