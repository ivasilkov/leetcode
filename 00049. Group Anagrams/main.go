package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println('9' - '0')
	fmt.Println(key("bat"))
	fmt.Println(key("atb"))
}

type node struct {
	next *node
	val  string
}
type list struct {
	head *node
	len  int
}

func (l *list) add(s string) {
	n := &node{
		next: l.head,
		val:  s,
	}
	l.head = n
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]*list)
	for i := range strs {
		k := key(strs[i])
		m[k].add(strs[i])
	}

	out := make([][]string, 0, len(m))
	for _, l := range m {
		tmp := make([]string, 0, l.len)
		for cur := l.head; cur != nil; cur = cur.next {
			tmp = append(tmp, cur.val)
		}
		out = append(out, tmp)
	}
	return out
}

func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for i := range strs {
		k := key(strs[i])
		m[k] = append(m[k], strs[i])
	}

	out := make([][]string, 0, len(m))
	for _, arr := range m {
		out = append(out, arr)
	}
	return out
}

func key(s string) string {
	k := []byte(s)
	sort.Slice(k, func(i, j int) bool {
		return k[i] > k[j]
	})
	return string(k)
}
