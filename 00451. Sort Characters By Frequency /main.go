package main

import (
	"fmt"
	"sort"
)

func main() {
	r := frequencySort("2a554442f544asfasssffffasss")
	fmt.Println(r)
}

func frequencySort(s string) string {
	arr := [62]uint16{} // 10+26+26
	for i := range s {
		arr[getKey(s[i])]++
	}
	out := []byte(s)
	sort.Slice(out, func(i, j int) bool {
		iKey, jKey := getKey(out[i]), getKey(out[j])
		if arr[iKey] == arr[jKey] {
			return out[i] > out[j]
		}
		return arr[iKey] > arr[jKey]
	})
	return string(out)
}

func getKey(b byte) byte {
	if b >= 'a' {
		return 10 + b - 'a'
	}
	if b >= 'A' {
		return 36 + b - 'A'
	}
	return b - '0'
}

////

func frequencySort2(s string) string {
	arr := [255]uint16{}
	for i := range s {
		arr[s[i]]++
	}

	out := []byte(s)
	sort.Slice(out, func(i, j int) bool {
		if arr[out[i]] == arr[out[j]] {
			return out[i] > out[j]
		}
		return arr[out[i]] > arr[out[j]]
	})
	return string(out)
}

/////

type item struct {
	val   byte
	count int
}

func frequencySort1(s string) string {
	m := map[byte]int{}
	var objs []item
	for i := range s {
		if idx, ok := m[s[i]]; ok {
			objs[idx].count++
		} else {
			m[s[i]] = len(objs)
			objs = append(objs, item{
				val:   s[i],
				count: 1,
			})
		}
	}
	sort.Slice(objs, func(i, j int) bool {
		return objs[i].count > objs[j].count
	})
	out := make([]byte, len(s))
	cur := 0
	for _, obj := range objs {
		for i := 0; i < obj.count; i++ {
			out[cur] = obj.val
			cur++
		}
	}
	return string(out)
}
