package _0094__Binary_Tree_Inorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
func inorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)
	traverse(root, &out)
	return out
}

func traverse(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, out)
	*out = append(*out, root.Val)
	traverse(root.Right, out)
}

// iterative
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	out := make([]int, 0)
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, v.Val)
		cur = v.Right
	}
	return out
}
