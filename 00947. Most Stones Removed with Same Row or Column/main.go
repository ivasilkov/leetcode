package main

import "fmt"

func main() {
	r := removeStones([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 3}, {2, 3}, {0, 2}})
	fmt.Println(r)
}

// Основная идея сводится к тому, что все связанные камни могут быть сведены к одному
// А значит после удаления останется камней = количество несвязанных групп
func removeStones(stones [][]int) int {
	cols, rows := map[int][]int{}, map[int][]int{}
	for i, stone := range stones {
		rows[stone[0]] = append(rows[stone[0]], i)
		cols[stone[1]] = append(cols[stone[1]], i)
	}

	visited := make([]bool, len(stones))
	groups := 0
	for i := range stones {
		if visited[i] {
			continue
		}
		groups++
		dfs(cols, rows, stones, visited, i)
	}
	return len(stones) - groups
}

func dfs(cols, rows map[int][]int, stones [][]int, visited []bool, idx int) {
	visited[idx] = true
	for _, i := range rows[stones[idx][0]] {
		if visited[i] {
			continue
		}
		dfs(cols, rows, stones, visited, i)
	}
	for _, i := range cols[stones[idx][1]] {
		if visited[i] {
			continue
		}
		dfs(cols, rows, stones, visited, i)
	}
}

///////////////////////////////////

//func removeStones(stones [][]int) int {
//	s := NewRemoveStonesSolution(stones)
//	return s.Solution()
//}
//
//type RemoveStonesSolution struct {
//	cols, rows map[int][]int
//	visited    []bool
//	stones     [][]int
//	groups     int
//}
//
//func NewRemoveStonesSolution(stones [][]int) *RemoveStonesSolution {
//	return &RemoveStonesSolution{
//		cols:    map[int][]int{},
//		rows:    map[int][]int{},
//		stones:  stones,
//		visited: make([]bool, len(stones)),
//	}
//}
//
//func (s *RemoveStonesSolution) Solution() int {
//	for i, stone := range s.stones {
//		s.rows[stone[0]] = append(s.rows[stone[0]], i)
//		s.cols[stone[1]] = append(s.rows[stone[1]], i)
//	}
//
//	for i := range s.stones {
//		if s.visited[i] {
//			continue
//		}
//		s.groups++
//		s.dfs(i)
//	}
//	return len(s.stones) - s.groups
//}
//
//func (s *RemoveStonesSolution) dfs(idx int) {
//	s.visited[idx] = true
//	for _, i := range s.rows[s.stones[idx][0]] {
//		if s.visited[i] {
//			continue
//		}
//		s.dfs(i)
//	}
//	for _, i := range s.cols[s.stones[idx][1]] {
//		if s.visited[i] {
//			continue
//		}
//		s.dfs(i)
//	}
//}
