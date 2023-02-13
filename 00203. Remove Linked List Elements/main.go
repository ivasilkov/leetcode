package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	ptr := head
	for ptr.Next != nil {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}
