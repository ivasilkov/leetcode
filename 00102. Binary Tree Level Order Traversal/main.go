package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	r := levelOrder(root)
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	out := [][]int{{}}

	lvl, nextLevelIdx := 0, 0
	for idx := 0; idx != len(queue); idx++ {
		v := queue[idx]
		out[lvl] = append(out[lvl], v.Val)
		if v.Left != nil {
			queue = append(queue, v.Left)
		}
		if v.Right != nil {
			queue = append(queue, v.Right)
		}

		if idx == nextLevelIdx && idx != len(queue)-1 {
			lvl++
			out = append(out, []int{})
			nextLevelIdx = len(queue) - 1
		}
	}

	return out
}
