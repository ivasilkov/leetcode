package _1339__Maximum_Product_of_Splitted_Binary_Tree_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ///////////
func maxProduct(root *TreeNode) int {
	total, result := sum(root), 0
	dfs(root, total, &result)
	return result % 1000000007
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sum(root.Left) + sum(root.Right)
}

func dfs(root *TreeNode, total int, result *int) int {
	if root == nil {
		return 0
	}

	sub := root.Val + dfs(root.Left, total, result) + dfs(root.Right, total, result)
	*result = max(*result, (total-sub)*sub)
	return sub
}

////////////

func maxProduct1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := map[*TreeNode]int{}
	total := sum1(root, m)

	return dfs1(root, m, total) % 1000000007
}

func sum1(root *TreeNode, m map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	m[root] = root.Val + sum1(root.Left, m) + sum1(root.Right, m)
	return m[root]
}

func dfs1(root *TreeNode, m map[*TreeNode]int, total int) int {
	if root == nil {
		return 0
	}
	left, right := 0, 0
	if root.Left != nil {
		l := m[root.Left]
		left = max((total-l)*l, dfs1(root.Left, m, total))
	}
	if root.Right != nil {
		r := m[root.Right]
		right = max((total-r)*r, dfs1(root.Right, m, total))
	}
	return max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
