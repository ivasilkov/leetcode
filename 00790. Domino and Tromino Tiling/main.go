package main

func main() {

}

func numTilings(n int) int {
	if n < 3 {
		return n
	}
	if n == 3 {
		return 5
	}
	pprev, prev, cur := 1, 2, 5
	for i := 3; i < n; i++ {
		tmp := (cur*2 + pprev) % 1000000007
		pprev, prev, cur = prev, cur, tmp
	}
	return cur
}

func numTilings1(n int) int {
	dp := make([]int, max(n, 3))
	dp[0], dp[1], dp[2] = 1, 2, 5
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1]*2 + dp[i-3]
		dp[i] = dp[i] % 1000000007
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
