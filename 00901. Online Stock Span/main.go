package main

import "fmt"

func main() {
	ss := Constructor()
	fmt.Println(ss.Next(100))
	fmt.Println(ss.Next(80))
	fmt.Println(ss.Next(60))
	fmt.Println(ss.Next(70))
	fmt.Println(ss.Next(60))
	fmt.Println(ss.Next(75))
	fmt.Println(ss.Next(85))
}

type stack []item

func (s *stack) push(i item) {
	*s = append(*s, i)
}
func (s *stack) pop() item {
	i := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return i
}
func (s *stack) peak() item {
	i := (*s)[len(*s)-1]
	return i
}

type item struct {
	span  int
	price int
}

type StockSpanner struct {
	arr [][2]int // 0 - span; 1 - price
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (ss *StockSpanner) Next(price int) int {
	out := 1
	for len(ss.arr) > 0 && ss.arr[len(ss.arr)-1][1] <= price {
		out += ss.arr[len(ss.arr)-1][0]
		ss.arr = ss.arr[:len(ss.arr)-1]
	}
	ss.arr = append(ss.arr, [2]int{out, price})
	return out
}
