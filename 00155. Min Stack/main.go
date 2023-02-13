package main

import "fmt"

func main() {
	s := Constructor()
	s.Push(2)
	s.Push(1)
	fmt.Println(s.Top(), s.GetMin())
	s.Pop()
	fmt.Println(s.Top(), s.GetMin())
}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}
func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
func (s *stack) peak() int {
	return (*s)[len(*s)-1]
}

type MinStack struct {
	data stack
	min  stack
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.data.push(val)
	if len(s.min) > 0 {
		if prevMin := s.min.peak(); prevMin < val {
			val = prevMin
		}
	}
	s.min.push(val)
}

func (s *MinStack) Pop() {
	s.data.pop()
	s.min.pop()
}

func (s *MinStack) Top() int {
	return s.data.peak()
}

func (s *MinStack) GetMin() int {
	return s.min.peak()
}

/////

type MinStack1 struct {
	arr    []int
	minIdx int
}

func Constructor1() MinStack1 {
	return MinStack1{}
}

func (s *MinStack1) Push(val int) {
	s.arr = append(s.arr, val)
	if s.arr[s.minIdx] > val {
		s.minIdx = len(s.arr) - 1
	}
}

func (s *MinStack1) Pop() {
	if s.minIdx == len(s.arr)-1 {
		s.minIdx = s.findMinIdx()
	}
	s.arr = s.arr[:len(s.arr)-1]
}

func (s *MinStack1) Top() int {
	return s.arr[len(s.arr)-1]
}

func (s *MinStack1) GetMin() int {
	return s.arr[s.minIdx]
}

func (s *MinStack1) findMinIdx() int {
	minIdx := 0
	for i := 1; i < len(s.arr)-1; i++ {
		if s.arr[i] < s.arr[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

/**
 * Your MinStack1 object will be instantiated and called as such:
 * obj := Constructor1();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
