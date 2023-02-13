package _0542__01_Matrix

type point struct {
	i, j int
}

func updateMatrix(mat [][]int) [][]int {
	q := []point{}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				q = append(q, point{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	directions := []point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for iter := 1; len(q) > 0; iter++ {
		for _, cur := range q {
			q = q[1:]
			for _, dir := range directions {
				x, y := cur.i-dir.i, cur.j-dir.j
				if x < 0 || y < 0 || x >= len(mat) || y >= len(mat[0]) || mat[x][y] != -1 {
					continue
				}
				mat[x][y] = iter
				q = append(q, point{x, y})
			}
		}
	}

	return mat
}
