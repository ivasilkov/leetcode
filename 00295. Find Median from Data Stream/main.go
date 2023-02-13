package main

import (
	"container/heap"
	"fmt"
)

func main() {
	f := Constructor()
	f.AddNum(6)
	fmt.Println(f.FindMedian())
	f.AddNum(10)
	fmt.Println(f.FindMedian())
	f.AddNum(2)
	fmt.Println(f.FindMedian())

	fmt.Println()
	//
	//sll := &sortedLinkedList{}
	//sll.push(1)
	//sll.push(6)
	//sll.push(2)
	//fmt.Println()
}

type Heap []int

func (h *Heap) Len() int           { return len(*h) }
func (h *Heap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *Heap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

type MedianFinder struct {
	min *Heap
	max *Heap
}

func Constructor() MedianFinder {
	min, max := &Heap{}, &Heap{}
	heap.Init(min)
	heap.Init(max)
	return MedianFinder{
		min: min,
		max: max,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.max.Len() == 0 || num > (*mf.max)[0] {
		heap.Push(mf.max, num)
	} else {
		heap.Push(mf.min, -num)
	}

	if mf.max.Len() > mf.min.Len()+1 {
		heap.Push(mf.min, -heap.Pop(mf.max).(int))
	} else if mf.min.Len() > mf.max.Len()+1 {
		heap.Push(mf.max, -heap.Pop(mf.min).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.max.Len() > mf.min.Len() {
		return float64((*mf.max)[0])
	}
	if mf.max.Len() < mf.min.Len() {
		return float64(-(*mf.min)[0])
	}

	return float64(-(*mf.min)[0]+(*mf.max)[0]) / 2
}

//////

type MedianFinder1 struct {
	min *sortedLinkedList
	max *sortedLinkedList
}

func Constructor1() *MedianFinder1 {
	return &MedianFinder1{
		min: &sortedLinkedList{},
		max: &sortedLinkedList{},
	}
}

func (mf *MedianFinder1) AddNum(num int) {
	if mf.max.len == 0 || num > mf.max.peak() {
		mf.max.push(num)
	} else {
		mf.min.push(-num)
	}

	if mf.max.len > mf.min.len+1 {
		mf.min.push(-mf.max.pop())
	} else if mf.min.len > mf.max.len+1 {
		mf.max.push(-mf.min.pop())
	}
}

func (mf *MedianFinder1) FindMedian() float64 {
	if mf.max.len > mf.min.len {
		return float64(mf.max.peak())
	}
	if mf.max.len < mf.min.len {
		return float64(-mf.min.peak())
	}

	return float64(-mf.min.peak()+mf.max.peak()) / 2
}

type node struct {
	next *node
	val  int
}

type sortedLinkedList struct {
	head *node
	len  int
}

func (l *sortedLinkedList) push(v int) {
	if l.head == nil || v < l.head.val {
		n := &node{
			next: l.head,
			val:  v,
		}
		l.head = n
		l.len++
		return
	}

	cur := l.head
	for ; cur.next != nil; cur = cur.next {
		if v > cur.next.val {
			continue
		}
		n := &node{
			next: cur.next,
			val:  v,
		}
		cur.next = n
		l.len++
		return
	}
	cur.next = &node{
		next: nil,
		val:  v,
	}
	l.len++
}

func (l *sortedLinkedList) peak() int {
	return l.head.val
}

func (l *sortedLinkedList) pop() int {
	v := l.head.val
	l.head = l.head.next
	l.len--
	return v
}
