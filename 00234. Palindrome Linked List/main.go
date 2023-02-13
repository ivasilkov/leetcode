package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						//Next: nil,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversed := reverseLinkedList(slow)
	for reversed != nil {
		if head.Val != reversed.Val {
			return false
		}
		head, reversed = head.Next, reversed.Next
	}
	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}
