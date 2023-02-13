package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	idx := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			idx = i
			break
		}
	}

	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+2:], inorder[idx+1:])
	return root
}
