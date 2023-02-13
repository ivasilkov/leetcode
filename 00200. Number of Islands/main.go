package _0200__Number_of_Islands

func numIslands(grid [][]byte) int {
	islands := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '1' {
				continue
			}
			dfs(grid, row, col)
			islands++
		}
	}
	return islands
}

func dfs(grid [][]byte, row, col int) {
	grid[row][col] = '0'
	// left
	if col > 0 {
		dfs(grid, row, col-1)
	}
	// right
	if col < len(grid[0])-1 {
		dfs(grid, row, col+1)
	}
	// top
	if row > 0 {
		dfs(grid, row-1, col)
	}
	// bot
	if row < len(grid)-1 {
		dfs(grid, row+1, col)
	}
}
