package _0104__Maximum_Depth_of_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// top - down

func maxDepth1(root *TreeNode) int {
	return traverse(root, 0)
}

func traverse(root *TreeNode, lvl int) int {
	if root == nil {
		return lvl
	}
	return max(traverse(root.Left, lvl+1), traverse(root.Right, lvl+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// down - top

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
