package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}

	c := Constructor()
	r := c.serialize(root)
	d := c.deserialize(r)
	fmt.Println(d)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	arr := []*TreeNode{root}
	for idx := 0; idx < len(arr); idx++ {
		v := arr[idx]
		if v == nil {
			continue
		}
		arr = append(arr, v.Left, v.Right)
	}

	out := make([]string, len(arr))
	for i := range arr {
		if arr[i] == nil {
			out[i] = "_"
		} else {
			out[i] = strconv.Itoa(arr[i].Val)
		}
	}
	return strings.Join(out, ".")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	arr := strings.Split(data, ".")
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(arr[0])

	q := []*TreeNode{root}
	for slow, fast := 0, 1; fast < len(arr); slow, fast = slow+1, fast+2 {
		n := q[slow]
		if v, err := strconv.Atoi(arr[fast]); err == nil {
			n.Left = &TreeNode{Val: v}
			q = append(q, n.Left)
		}
		if v, err := strconv.Atoi(arr[fast+1]); err == nil {
			n.Right = &TreeNode{Val: v}
			q = append(q, n.Right)
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
