package _1026__Maximum_Difference_Between_Node_and_Ancestor_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, mn, mx int) int {
	if root == nil {
		return mx - mn
	}
	mn = min(root.Val, mn)
	mx = max(root.Val, mx)
	return max(
		dfs(root.Left, mn, mx),
		dfs(root.Right, mn, mx),
	)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
