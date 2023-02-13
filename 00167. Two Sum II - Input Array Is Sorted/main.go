package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		}

		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{-1, -1}
}
