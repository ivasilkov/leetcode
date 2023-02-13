package _0112__Path_Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	left := hasPathSum(root.Left, targetSum-root.Val)
	right := hasPathSum(root.Right, targetSum-root.Val)
	return left || right
}

////

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return check(root, 0, targetSum)
}

func check(root *TreeNode, sum, target int) bool {
	if root.Left == nil && root.Right == nil {
		return sum+root.Val == target
	}
	result := false
	if root.Left != nil {
		result = result || check(root.Left, sum+root.Val, target)
	}
	if root.Right != nil {
		result = result || check(root.Right, sum+root.Val, target)
	}
	return result
}
