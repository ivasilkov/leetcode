package main

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
					Val: 4,
					Next: &ListNode{
						Val: 5,
						//Next: nil,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	oddEvenList(head)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	cur := head
	evenHead := head.Next
	evenCur := head.Next
	for evenCur != nil && evenCur.Next != nil {
		cur.Next = cur.Next.Next
		evenCur.Next = evenCur.Next.Next
		evenCur = evenCur.Next
		cur = cur.Next
	}
	cur.Next = evenHead
	return head
}

func oddEvenList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slowPtr, fastPtr := head, head
	for fastPtr.Next != nil {
		if fastPtr.Next.Val%2 != head.Val%2 {
			fastPtr = fastPtr.Next
			continue
		}
		slowPtr.Next, fastPtr.Next = fastPtr.Next, slowPtr.Next
		slowPtr.Next.Next, fastPtr.Next.Next = fastPtr.Next.Next, slowPtr.Next.Next

		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}
	return head
}
