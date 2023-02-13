package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(2))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.GetRandom())
}

type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{}, arr: []int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.m[val]; ok {
		return false
	}
	l := len(rs.arr)
	rs.arr = append(rs.arr, val)
	rs.m[val] = l
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.m[val]
	if !ok {
		return false
	}

	// move from the end of slice to the index of found value
	rs.arr[idx] = rs.arr[len(rs.arr)-1]
	// update index in map for moved value
	rs.m[rs.arr[idx]] = idx
	// remove last element from array
	rs.arr = rs.arr[:len(rs.arr)-1]
	// remove key from map
	delete(rs.m, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(rs.arr))
	return rs.arr[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
