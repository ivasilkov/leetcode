package _0145__Binary_Tree_Postorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/// recursive

func postorderTraversal(root *TreeNode) []int {
	var out []int
	traverse(root, &out)
	return out
}

func traverse(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, out)
	traverse(root.Right, out)
	*out = append(*out, root.Val)
}

/// iterative

func postorderTraversal1(root *TreeNode) []int {
	var out []int
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			out = append([]int{cur.Val}, out...)
			cur = cur.Right
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = v.Left
		}
	}

	return out
}
