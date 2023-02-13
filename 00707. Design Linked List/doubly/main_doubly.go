package main

func main() {
	l := Constructor()
	l.AddAtIndex(0, 10)
}

type Node struct {
	next *Node
	prev *Node
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
	if index >= l.len {
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
		prev: nil,
		val:  val,
	}
	if l.head != nil {
		l.head.prev = n
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
		next: nil,
		prev: last,
		val:  val,
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
	if index == l.len {
		l.AddAtTail(val)
		return
	}

	node := l.head
	for i := 1; i < index; i++ {
		node = node.next
	}
	next := node.next
	n := &Node{
		next: next,
		prev: node,
		val:  val,
	}
	node.next = n
	next.prev = n
	l.len++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.len {
		return
	}
	defer func() { l.len-- }()
	if index == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		return
	}

	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	node.prev.next = node.next
	if node.next != nil {
		node.next.prev = node.prev
	}
}
