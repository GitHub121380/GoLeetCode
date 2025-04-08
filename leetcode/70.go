package main

/*
leetcode 70
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

// 递归
func climbStairs(n int) int {
	//if n == 1 {
	//	return 1
	//}
	//if n == 2 {
	//	return 2
	//}
	//return climbStairs(n-1) + climbStairs(n-2)

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
