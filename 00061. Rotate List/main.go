package main

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	rotateRight(h, 1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	fast := head
	for i := 0; i < k; i++ {
		if fast.Next == nil {
			return rotateRight(head, k%(i+1))
		}
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	l, last := 1, head
	for last.Next != nil {
		l, last = l+1, last.Next
	}

	k = k % l
	if k == 0 {
		return head
	}

	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	last.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}
