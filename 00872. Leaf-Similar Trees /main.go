package _0872__Leaf_Similar_Trees_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := &[]*TreeNode{root1}, &[]*TreeNode{root2}
	for len(*s1) != 0 && len(*s2) != 0 {
		if dfs(s1) != dfs(s2) {
			return false
		}
	}
	return len(*s1) == 0 && len(*s2) == 0
}

func dfs(s *[]*TreeNode) int {
	for {
		node := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		if node.Right != nil {
			*s = append(*s, node.Right)
		}
		if node.Left != nil {
			*s = append(*s, node.Left)
		}
		if node.Left == nil && node.Right == nil {
			return node.Val
		}
	}
}

////////////

func leafSimilar1(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := &[]*TreeNode{}
	leafs(root1, leafs1)

	leafs2 := &[]*TreeNode{}
	leafs(root2, leafs2)

	if len(*leafs1) != len(*leafs2) {
		return false
	}

	for i := range *leafs1 {
		if (*leafs1)[i].Val != (*leafs2)[i].Val {
			return false
		}
	}
	return true
}

func leafs(root *TreeNode, out *[]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*out = append(*out, root)
	}
	leafs(root.Left, out)
	leafs(root.Right, out)
}
