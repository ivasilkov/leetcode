package main

import "fmt"

func main() {
	r := buildTree([]int{1, 2}, []int{1, 2})
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	idx := len(inorder) - 1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[i] {
			continue
		}
		idx = i
		break
	}

	root := &TreeNode{Val: inorder[idx]}
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}
