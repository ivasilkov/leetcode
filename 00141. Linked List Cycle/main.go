package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	f := &ListNode{
		Val: -4,
	}
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  0,
				Next: f,
			},
		},
	}
	f.Next = head.Next

	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
