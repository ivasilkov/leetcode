package main

import "fmt"

func main() {
	arr := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	entrance := []int{1, 2}

	//arr := [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}
	//entrance := []int{1, 0}

	r := nearestExit(arr, entrance)
	fmt.Println(r)
}

func nearestExit(maze [][]byte, entrance []int) int {
	q := [][2]int{{entrance[0], entrance[1]}}
	m, n := len(maze), len(maze[0])
	maze[entrance[0]][entrance[1]] = '+'

	steps := 0
	for len(q) > 0 {
		for _, v := range q {
			q = q[1:]
			if (v[0] == m-1 || v[1] == n-1 || v[0] == 0 || v[1] == 0) && !(v[0] == entrance[0] && v[1] == entrance[1]) {
				return steps
			}
			for _, div := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				x, y := v[0]+div[0], v[1]+div[1]
				if x >= 0 && y >= 0 && x < m && y < n && maze[x][y] != '+' {
					q = append(q, [2]int{x, y})
					maze[x][y] = '+'
				}
			}
		}
		steps++
	}
	return -1
}

func nearestExit1(maze [][]byte, entrance []int) int {
	entranceIdx := intsToIdx(entrance[0], entrance[1])
	q, visited := []int{entranceIdx}, map[int]struct{}{}
	m, n := len(maze), len(maze[0])

	steps := -1
	for len(q) > 0 {
		steps++
		for _, v := range q {
			q = q[1:] // remove first element
			if _, ok := visited[v]; ok {
				continue
			}
			visited[v] = struct{}{}

			i, j := fromIdxToInt(v)
			if maze[i][j] == '.' && (i == m-1 || j == n-1 || i == 0 || j == 0) && v != entranceIdx {
				return steps
			}
			if maze[i][j] == '+' {
				continue
			}
			if i > 0 {
				q = append(q, intsToIdx(i-1, j))
			}
			if i < len(maze)-1 {
				q = append(q, intsToIdx(i+1, j))
			}
			if j > 0 {
				q = append(q, intsToIdx(i, j-1))
			}
			if j < len(maze[0])-1 {
				q = append(q, intsToIdx(i, j+1))
			}
		}
	}
	return -1
}

func intsToIdx(i, j int) int {
	return i*100 + j // 1 <= m,n <= 100; then max index == 99
}

func fromIdxToInt(idx int) (int, int) {
	return idx / 100, idx % 100
}
