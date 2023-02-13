package main

import "fmt"

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

type MyQueue struct {
	direct   stack
	reversed stack
}

func Constructor() MyQueue {
	return MyQueue{
		direct:   stack{},
		reversed: stack{},
	}
}

func (q *MyQueue) Push(x int) {
	q.direct.push(x)
}

func (q *MyQueue) Pop() int {
	if q.reversed.empty() {
		q.move()
	}

	return q.reversed.pop()
}

func (q *MyQueue) Peek() int {
	if q.reversed.empty() {
		q.move()
	}

	return q.reversed.peak()
}

func (q *MyQueue) Empty() bool {
	return q.direct.empty() && q.reversed.empty()
}

func (q *MyQueue) move() {
	for !q.direct.empty() {
		q.reversed.push(q.direct.pop())
	}
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
	v := (*s)[len(*s)-1]
	return v
}
func (s *stack) empty() bool {
	return len(*s) == 0
}
