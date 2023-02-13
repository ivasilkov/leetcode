package main

func main() {
	s := Constructor()
	s.Add(2)
	s.Add(22)
	s.Add(32)
	s.Add(42)
	s.Remove(2)
	s.Remove(22)
	s.Remove(32)
	s.Remove(42)
	//s.Remove(22)
	//s.Contains(22)
	//s.Add(1)
	//s.Add(2)
	//fmt.Println(s.Contains(1))
	//fmt.Println(s.Contains(3))
	//s.Add(2)
	//s.Add(22)
	//fmt.Println(s.Contains(2))
	//s.Remove(22)
	//fmt.Println(s.Contains(2))
	//fmt.Println()
}

type Node struct {
	next *Node
	val  int
}

type MyHashSet struct {
	data []*Node
}

const size = 8

func Constructor() MyHashSet {
	return MyHashSet{
		data: make([]*Node, size),
	}
}

func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	idx := s.idx(key)
	n := &Node{
		next: s.data[idx],
		val:  key,
	}
	s.data[idx] = n
}

func (s *MyHashSet) Remove(key int) {
	idx := s.idx(key)
	if s.data[idx] == nil {
		return
	}
	if s.data[idx].val == key {
		s.data[idx] = s.data[idx].next
		return
	}

	for cur := s.data[idx]; cur.next != nil; cur = cur.next {
		if cur.next.val != key {
			continue
		}
		cur.next = cur.next.next
		break
	}
}

func (s *MyHashSet) Contains(key int) bool {
	idx := s.idx(key)
	for cur := s.data[idx]; cur != nil; cur = cur.next {
		if cur.val == key {
			return true
		}
	}
	return false
}

func (s *MyHashSet) idx(key int) int {
	return key % size
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
