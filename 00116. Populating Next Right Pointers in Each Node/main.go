package _0116__Populating_Next_Right_Pointers_in_Each_Node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Right == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	connect(root.Left)
	connect(root.Right)
	return root
}
