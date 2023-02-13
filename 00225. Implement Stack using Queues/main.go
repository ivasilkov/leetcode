package main

import "fmt"

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(4)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}

type MyStack struct {
	out queue
}

func Constructor() MyStack {
	return MyStack{
		out: queue{},
	}
}

func (s *MyStack) Push(x int) {
	tmp := queue{x}
	for !s.out.empty() {
		tmp.enqueue(s.out.dequeue())
	}
	s.out = tmp
}

func (s *MyStack) Pop() int {
	return s.out.dequeue()
}

func (s *MyStack) Top() int {
	return s.out.top()
}

func (s *MyStack) Empty() bool {
	return s.out.empty()
}

type queue []int

func (q *queue) enqueue(v int) {
	*q = append(*q, v)
}
func (q *queue) dequeue() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}
func (q *queue) top() int {
	return (*q)[0]
}
func (q *queue) empty() bool {
	return len(*q) == 0
}
