package main

import "fmt"

func main() {
	m := Constructor()
	m.Put(1, 1)
	m.Put(9, 0)
	m.Remove(1)
	m.Remove(9)
	//m.Put(2, 2)
	//fmt.Println(m.Get(1))
	//fmt.Println(m.Get(3))
	//m.Put(2, 1)
	//fmt.Println(m.Get(2))
	//m.Remove(2)
	//fmt.Println(m.Get(2))
	fmt.Println()
}

type node struct {
	next *node
	val  int
	key  int
}

type MyHashMap struct {
	data []*node
}

const size = 8

func Constructor() MyHashMap {
	return MyHashMap{data: make([]*node, 8)}
}

func (m *MyHashMap) Put(key int, value int) {
	idx := m.idx(key)
	if m.data[idx] == nil {
		n := &node{
			next: m.data[idx],
			val:  value,
			key:  key,
		}
		m.data[idx] = n
		return
	}

	for cur := m.data[idx]; cur != nil; cur = cur.next {
		if cur.key == key {
			cur.val = value
			return
		}
	}

	n := &node{
		next: m.data[idx],
		val:  value,
		key:  key,
	}
	m.data[idx] = n
}

func (m *MyHashMap) Get(key int) int {
	idx := m.idx(key)
	for cur := m.data[idx]; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur.val
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	idx := m.idx(key)
	if m.data[idx] == nil {
		return
	}
	if m.data[idx].key == key {
		m.data[idx] = m.data[idx].next
		return
	}

	for cur := m.data[idx]; cur.next != nil; cur = cur.next {
		if cur.next.key == key {
			cur.next = cur.next.next
			break
		}
	}
}

func (m *MyHashMap) idx(key int) int {
	return key % size
}
