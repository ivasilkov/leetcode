package _0117__Populating_Next_Right_Pointers_in_Each_Node_II_

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || (root.Right == nil && root.Left == nil) {
		return root
	}

	if root.Left == nil {
		root.Right.Next = nextChild(root.Next)
	} else if root.Right == nil {
		root.Left.Next = nextChild(root.Next)
	} else {
		root.Left.Next = root.Right
		root.Right.Next = nextChild(root.Next)
	}

	connect(root.Right)
	connect(root.Left)
	return root
}

func nextChild(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return nextChild(root.Next)
	}
	if root.Left != nil {
		return root.Left
	}
	return root.Right
}
