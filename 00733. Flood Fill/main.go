package main

import "fmt"

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	r := floodFill(image, 1, 1, 2)
	fmt.Println(r)
}

func floodFill1(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	dfs(image, sr, sc, image[sr][sc], color)
	return image
}

func dfs(image [][]int, i int, j int, startColor, newColor int) {
	if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) {
		return
	}
	if image[i][j] != startColor {
		return
	}
	image[i][j] = newColor
	dfs(image, i+1, j, startColor, newColor)
	dfs(image, i-1, j, startColor, newColor)
	dfs(image, i, j+1, startColor, newColor)
	dfs(image, i, j-1, startColor, newColor)
}

///

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	startColor := image[sr][sc]
	q := [][2]int{{sr, sc}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		image[v[0]][v[1]] = color
		if v[0] > 0 && image[v[0]-1][v[1]] != startColor {
			q = append(q, [2]int{v[0] - 1, v[1]})
		}
		if v[0] < len(image)-1 && image[v[0]+1][v[1]] != startColor {
			q = append(q, [2]int{v[0] + 1, v[1]})
		}
		if v[1] > 0 && image[v[0]][v[1]-1] != startColor {
			q = append(q, [2]int{v[0], v[1] - 1})
		}
		if v[1] < len(image[0])-1 && image[v[0]][v[1]+1] != startColor {
			q = append(q, [2]int{v[0], v[1] + 1})
		}
	}

	return image
}
