package main

import "fmt"

func main() {
	arr := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}
	out := uniquePathsIII(arr)
	fmt.Println(out)
}

func uniquePathsIII(grid [][]int) int {
	empty, n, m := 1, len(grid), len(grid[0])
	startX, startY := 0, 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				startX, startY = i, j
			} else if grid[i][j] == 0 {
				empty++
			}
		}
	}

	out := 0
	var dfs func(x int, y int, count int)
	dfs = func(x int, y int, count int) {
		if x == n || y == m || x < 0 || y < 0 || grid[x][y] == -1 {
			return
		}
		if grid[x][y] == 2 {
			if empty == count {
				out++
			}
			return
		}
		grid[x][y] = -1
		dfs(x+1, y, count+1)
		dfs(x-1, y, count+1)
		dfs(x, y+1, count+1)
		dfs(x, y-1, count+1)
		grid[x][y] = 0
	}
	dfs(startX, startY, 0)
	return out
}
