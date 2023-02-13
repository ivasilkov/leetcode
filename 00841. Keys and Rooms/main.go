package _0841__Keys_and_Rooms

func canVisitAllRooms1(rooms [][]int) bool {
	visited := map[int]struct{}{}
	q := []int{0}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		visited[v] = struct{}{}
		for _, k := range rooms[v] {
			if _, ok := visited[k]; ok {
				continue
			}
			q = append(q, k)
		}
	}

	return len(visited) == len(rooms)
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	q := []int{0}
	locked := len(rooms) - 1

	for len(q) > 0 && locked != 0 {
		v := q[0]
		q = q[1:]

		for _, k := range rooms[v] {
			if visited[k] {
				continue
			}
			q = append(q, k)
			visited[k] = true
			locked--
		}
	}

	return locked == 0
}
