package main

import "fmt"

func main() {
	n2 := &Node{
		Val:  2,
		Next: nil,
	}
	n1 := &Node{
		Val:    1,
		Next:   n2,
		Random: n2,
	}
	n2.Random = n1
	r := copyRandomList(n1)
	fmt.Println(r)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// O(1) space
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{
			Val:    cur.Val,
			Next:   cur.Next,
			Random: cur.Random,
		}
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Next.Random != nil {
			cur.Next.Random = cur.Next.Random.Next
		}
	}

	out := head.Next
	for cur := head; cur != nil && cur.Next != nil; cur = cur.Next {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		if tmp.Next != nil {
			tmp.Next = tmp.Next.Next
		}
	}

	return out
}

// using map
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := map[*Node]*Node{}

	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}

	for cur := head; cur != nil; cur = cur.Next {
		m[cur].Random = m[cur.Random]
		m[cur].Next = m[cur.Next]
	}

	return m[head]
}
