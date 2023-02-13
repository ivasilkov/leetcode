package main

func main() {

}

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return 1
	}

	out := 0
	for i := 0; i < n; i++ {
		m := map[float64]int{}
		for j := 0; j < n; j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			slope := float64(y2-y1) / float64(x2-x1)
			m[slope]++
		}
		for _, v := range m {
			v++
			if out < v {
				out = v
			}
		}
	}
	return out
}
