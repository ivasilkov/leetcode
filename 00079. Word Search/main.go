package main

import "fmt"

func main() {
	r := exist([][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'e', 's'},
		{'a', 'd', 'e', 'e'},
	}, "abceseeefs")
	fmt.Println(r)
}

// bfs
func exist1(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	q := make([][2]int, m*n)

	for i := range board {
		for j := range board[i] {
			q = append(q, [2]int{i, j})
		}
	}

	for iter := 0; len(q) > 0 && iter < len(word); iter++ {
		for _, v := range q {
			q = q[1:]
			if board[v[0]][v[1]] != word[iter] {
				continue
			}
			if iter == len(word)-1 {
				return true
			}

			for _, diff := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				x, y := v[0]-diff[0], v[1]-diff[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					q = append(q, [2]int{x, y})
				}
			}
		}
	}

	return false
}

// dfs
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if board[i][j] != word[0] {
				continue
			}
			if dfs(i, j, 0, board, word) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j, pos int, board [][]byte, word string) bool {
	if pos == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if word[pos] != board[i][j] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = '-'

	out := dfs(i+1, j, pos+1, board, word) ||
		dfs(i-1, j, pos+1, board, word) ||
		dfs(i, j+1, pos+1, board, word) ||
		dfs(i, j-1, pos+1, board, word)

	board[i][j] = tmp

	return out
}
