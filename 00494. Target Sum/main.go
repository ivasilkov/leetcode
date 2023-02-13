package _0494__Target_Sum

func findTargetSumWays(nums []int, target int) int {
	return dfs(nums, target, 0)
}

func dfs(nums []int, target int, idx int) int {
	if idx == len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}

	plus := dfs(nums, target+nums[idx], idx+1)
	minus := dfs(nums, target-nums[idx], idx+1)

	return plus + minus
}
