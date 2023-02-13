package main

import "fmt"

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	r := cloneGraph(n1)
	fmt.Println(r)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(n *Node) *Node {
	m := map[int]*Node{}
	return dfs(m, n)
}

func dfs(m map[int]*Node, n *Node) *Node {
	if n == nil {
		return nil
	}

	clone := &Node{
		Val:       n.Val,
		Neighbors: make([]*Node, len(n.Neighbors)),
	}
	m[n.Val] = clone

	for i, neighbor := range n.Neighbors {
		if v, ok := m[neighbor.Val]; ok {
			clone.Neighbors[i] = v
		} else {
			clone.Neighbors[i] = dfs(m, n.Neighbors[i])
		}
	}

	return clone
}
