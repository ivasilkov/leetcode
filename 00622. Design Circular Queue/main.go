package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	q := Constructor(3)
	fmt.Println(q.Rear())
	q.EnQueue(1)
	fmt.Println(q.Rear())
	q.EnQueue(2)
	fmt.Println(q.Rear())
	q.EnQueue(3)
	fmt.Println(q.Rear())
	q.EnQueue(4)
	fmt.Println(q.Rear())
	fmt.Println(q.Front())
	q.DeQueue()
	fmt.Println(q.Front())
	q.DeQueue()
	fmt.Println(q.Front())
	q.DeQueue()
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	fmt.Println()
}

type MyCircularQueue struct {
	arr        []int
	head, tail int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{arr: make([]int, k), tail: -1}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.tail++
	q.arr[q.tail%len(q.arr)] = value
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head++
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.arr[q.head%len(q.arr)]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.arr[q.tail%len(q.arr)]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.head > q.tail
}

func (q *MyCircularQueue) IsFull() bool {
	return q.tail-q.head == len(q.arr)-1
}

/// generic

func ConstructorGen(k int) CircularQueue[int] {
	return *NewCircularQueue[int](k, -1)
}

type CircularQueue[T constraints.Ordered] struct {
	arr        []T
	defValue   T
	head, tail int
}

func NewCircularQueue[T constraints.Ordered](k int, defValue T) *CircularQueue[T] {
	return &CircularQueue[T]{arr: make([]T, k), tail: -1, defValue: defValue}
}

func (q *CircularQueue[T]) EnQueue(value T) bool {
	if q.IsFull() {
		return false
	}
	q.tail++
	q.arr[q.tail%len(q.arr)] = value
	return true
}

func (q *CircularQueue[T]) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head++
	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.IsEmpty() {
		return q.defValue
	}
	return q.arr[q.head%len(q.arr)]
}

func (q *CircularQueue[T]) Rear() T {
	if q.IsEmpty() {
		return q.defValue
	}
	return q.arr[q.tail%len(q.arr)]
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.head > q.tail
}

func (q *CircularQueue[T]) IsFull() bool {
	return q.tail-q.head == len(q.arr)-1
}

/**
 * Your CircularQueue object will be instantiated and called as such:
 * obj := NewCircularQueue(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
