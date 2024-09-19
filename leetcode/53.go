package main

// 贪心和动态规划其实公式都是一样的，都在于元素自己想跟一个正数加（这就是贪），那么究竟是加之前的最大和，还是选择自己重新开始，要看之前是正数还是负数
// 用dp[i]表示，以第 i 个数作为子数组结尾的「连续子数组的最大和」，状态转移方程是 dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)，dp[i] = nums[i] (dp[i-1] ≤ 0)。
// 指针指到的每个元素自己都想跟一个正的数相加而不是负数
// /////获取对应子序列的方法
// 所以当一个dp[i]从负变正，一定表示他是一个新序列的开始。而dp[i]本身又代表一个子序列的结束，所以当我们得到max(dp[i])后，往前推找到第一个从负变正的数，就能获得这个子序列的起始下标和终止下标了
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > max {
			max = sum
		}
		//sum = result[i-1] + nums[i]
		//if result[i-1] < 0 {
		//	result = append(result, nums[i])
		//} else {
		//	result = append(result, sum)
		//}
	}
	return max
}
