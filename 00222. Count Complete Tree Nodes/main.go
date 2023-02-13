package _0222__Count_Complete_Tree_Nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh, rh := height(root.Left), height(root.Right)
	if lh == rh {
		return 1<<lh + countNodes1(root.Right)
	}
	return 1<<rh + countNodes1(root.Left)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + height(root.Left)
}
