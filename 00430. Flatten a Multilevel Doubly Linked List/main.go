package main

import "fmt"

func main() {

	n3 := &Node{
		Val:   3,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	n2 := &Node{
		Val:   2,
		Prev:  nil,
		Next:  nil,
		Child: n3,
	}
	n1 := &Node{
		Val:   1,
		Prev:  nil,
		Next:  n2,
		Child: nil,
	}
	n2.Prev = n1
	r := flatten(n1)
	fmt.Println(r)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	cur := root
	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
			continue
		}

		right := cur.Child
		for right.Next != nil {
			right = right.Next
		}

		next := cur.Next
		cur.Next = cur.Child
		cur.Child.Prev = cur
		if next != nil {
			next.Prev = right
		}
		right.Next = next
		cur.Child = nil
		cur = cur.Next
	}

	return root
}

func flatten1(root *Node) *Node {
	var stack []*Node
	rPtr, wPtr := root, root
	for rPtr != nil {
		if rPtr.Child != nil {
			if rPtr.Next != nil {
				stack = append(stack, rPtr.Next)
			}
			rPtr, rPtr.Child = rPtr.Child, nil
		} else {
			if rPtr.Next == nil {
				if len(stack) > 0 {
					v := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					rPtr = v
				} else {
					rPtr = nil
				}
			} else {
				rPtr = rPtr.Next
			}
		}
		wPtr.Next = rPtr
		if rPtr != nil {
			rPtr.Prev = wPtr
		}
		wPtr = wPtr.Next
	}

	return root
}
