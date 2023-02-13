package _0070__Climbing_Stairs

func climbStairs(n int) int {
	prev, cur := 1, 1
	for i := 1; i < n; i++ {
		prev, cur = cur, cur+prev
	}
	return cur
}
