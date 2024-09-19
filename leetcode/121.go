package main

import "fmt"

// 买股票的思想 如果我是在当天之前的某个历史最低点买入，那我当天卖出一定是赚的最多的
// 在题目中，我们只要用一个变量记录一个历史最低价格 minprice，我们就可以假设自己的股票是在那天买的。那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice。
// 因此，我们只需要遍历价格数组一遍，记录历史最低点，然后在每一天考虑这么一个问题：如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？当考虑完所有天数之时，我们就得到了最好的答案。
// 并不是求全局历史最低点，而是每次都假设是今天卖出，然后求今天之前的历史最低点。而这个历史最低点并不需要额外遍历，而是每天考虑的时候顺带记录的。因此时间复杂度还是O(N)而不是O(N^2)。
func maxProfit(prices []int) int {
	result := 0
	min := 10000
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		d := prices[i] - min
		if d > result {
			result = d
		}
	}

	return result
}

//dp
//https://leetcode.cn/circle/discuss/qiAgHn/
/*
因此对 T[i][k] 的定义需要分成两项：
T[i][k][0] 表示在第 i 天结束时，最多进行 k 次交易且在进行操作后持有 0 份股票的情况下可以获得的最大收益；
T[i][k][1] 表示在第 i 天结束时，最多进行 k 次交易且在进行操作后持有 1 份股票的情况下可以获得的最大收益。
使用新的状态表示之后，可以得到基准情况和状态转移方程。
基准情况：
T[-1][k][0] = 0, T[-1][k][1] = -Infinity
T[i][0][0] = 0, T[i][0][1] = -Infinity
状态转移方程：
T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])
T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k - 1][0] - prices[i])
*/

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
