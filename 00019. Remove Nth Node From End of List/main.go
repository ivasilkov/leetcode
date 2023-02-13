package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	fastPtr := head
	for n > 0 {
		fastPtr, n = fastPtr.Next, n-1
	}
	if fastPtr == nil {
		return head.Next
	}
	slowPtr := head
	for fastPtr.Next != nil {
		slowPtr, fastPtr = slowPtr.Next, fastPtr.Next
	}
	slowPtr.Next = fastPtr.Next
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr, l := head, 0
	for ptr.Next != nil {
		ptr, l = ptr.Next, l+1
	}

	var prev *ListNode
	cur := head
	for i := 0; i < l-n; i++ {
		prev = cur
		cur = cur.Next
	}

	prev.Next = cur.Next
	return head
}
