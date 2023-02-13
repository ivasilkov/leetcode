package _0652__Find_Duplicate_Subtrees

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := map[string]uint8{}
	out := make([]*TreeNode, 0)
	key(m, &out, root)
	return out
}

func key(m map[string]uint8, out *[]*TreeNode, n *TreeNode) string {
	if n == nil {
		return "_"
	}
	leftKey := key(m, out, n.Left)
	rightKey := key(m, out, n.Right)
	k := fmt.Sprintf("(%s)(%d)(%s)", leftKey, n.Val, rightKey)
	m[k]++
	if m[k] == 2 {
		*out = append(*out, n)
	}
	return k
}
