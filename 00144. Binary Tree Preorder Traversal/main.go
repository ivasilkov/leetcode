package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack []*TreeNode

func (s *stack) push(i *TreeNode) {
	*s = append(*s, i)
}
func (s *stack) pop() *TreeNode {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	out := make([]int, 0)
	s := make(stack, 0)
	s.push(root)

	for len(s) != 0 {
		node := s.pop()
		if node == nil {
			continue
		}
		out = append(out, node.Val)
		s.push(node.Right)
		s.push(node.Left)
	}
	return out
}

func preorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)
	traverse(root, &out)
	return out
}

func traverse(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}
	*out = append(*out, root.Val)
	traverse(root.Left, out)
	traverse(root.Right, out)
}
