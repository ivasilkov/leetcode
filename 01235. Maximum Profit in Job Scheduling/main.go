package _1235__Maximum_Profit_in_Job_Scheduling

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([][3]int, len(startTime))
	for i := range startTime {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	memo := make([]int, len(jobs))
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i == len(jobs) {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		l, r := i+1, len(jobs)
		for l < r {
			mid := l + (r-l)/2
			if jobs[mid][0] < jobs[i][1] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		memo[i] = max(dfs(i+i), dfs(l)+jobs[i][2])
		return memo[i]
	}
	return dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
