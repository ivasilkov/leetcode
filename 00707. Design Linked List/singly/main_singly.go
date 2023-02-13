package main

import "fmt"

func main() {
	obj := Constructor()
	//param_1 := obj.Get(0)
	//fmt.Println(param_1)
	//obj.AddAtHead(1)
	obj.AddAtTail(3)
	//obj.AddAtIndex(0, 10)
	//obj.AddAtIndex(0, 20)
	//obj.AddAtIndex(0, 30)
	//obj.DeleteAtIndex(0)
	fmt.Println(obj.Get(0))
}

type Node struct {
	next *Node
	val  int
}

type MyLinkedList struct {
	head *Node
	len  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if l.len <= index {
		return -1
	}
	out := l.head
	for i := 0; i < index; i++ {
		out = out.next
	}
	return out.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	n := &Node{
		next: l.head,
		val:  val,
	}
	l.head = n
	l.len++
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.head == nil {
		l.AddAtHead(val)
		return
	}

	last := l.head
	for last.next != nil {
		last = last.next
	}

	n := &Node{
		val: val,
	}
	last.next = n
	l.len++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.len {
		return
	}
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	prev := l.head
	for i := 1; i < index; i++ {
		prev = prev.next
	}
	n := &Node{
		next: prev.next,
		val:  val,
	}
	prev.next = n
	l.len++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if l.len == 0 || index >= l.len {
		return
	}
	if index == 0 {
		l.head = l.head.next
		l.len--
		return
	}
	prev := l.head
	for i := 1; i < index; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	l.len--
}
