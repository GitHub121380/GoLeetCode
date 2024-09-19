package main

import "fmt"

// 动态规划
// 对于一个子串而言，如果它是回文串，并且长度大于 2，那么将它首尾的两个字母去除之后，它仍然是个回文串。例如对于字符串 “ababa”，如果我们已经知道 “bab” 是回文串，那么 “ababa” 一定是回文串，这是因为它的首尾两个字母都是 “a”。
// dp[i][j] = dp[i+1][j-1] && s[i]==s[j]  i-j的子串是否是回文，取决于两端字符是否相等以及去掉两端字符的子串是否回文
func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	dp := make([][]bool, 0, l)
	for i := 0; i < l; i++ {
		dp = append(dp, make([]bool, l))
	}
	//长度=1、2的情况
	for i := 0; i < l; i++ {
		dp[i][i] = true
		if i != l-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	//长度>=3的情况
	for i := l - 3; i >= 0; i-- {
		for j := i + 2; j < l; j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		}
	}
	left := 0
	right := 0
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if dp[i][j] && j-i > right-left {
				left, right = i, j
			}
		}
	}
	return s[left : right+1]
}

// 观察动态规划公式，发现dp[i][j]一定会依赖于dp[i+1][j-1]，那么一定会依赖于dp[i+2][j-2]
// 可以发现，所有的状态在转移的时候的可能性都是唯一的。也就是说，我们可以从每一种边界情况开始向两边「扩展」，也可以得出所有的状态对应的答案。
// 边界情况（中心情况）即为子串长度为 1 或 2 的情况。我们枚举每一种边界情况，并从对应的子串开始不断地向两边扩展。如果两边的字母相同，我们就可以继续扩展。
// 例如从 P(i+1,j−1) 扩展到 P(i,j)；如果两边的字母不同，我们就可以停止扩展，因为在这之后的子串不可能是回文串了。
// 聪明的读者此时应该可以发现，「边界情况」对应的子串实际上就是我们「扩展」出的回文串的「回文中心」。
// 方法二的本质即为：我们枚举所有的「回文中心」并尝试「扩展」，直到无法扩展为止，此时的回文串长度即为此「回文中心」下的最长回文串长度。我们对所有的长度求出最大值，即可得到最终的答案。
func longestPalindrome2(s string) string {
	if s == "" {
		return s
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := i, i
		for ; l1 >= 0 && r1 < len(s) && s[l1] == s[r1]; l1, r1 = l1-1, r1+1 {
		}
		l1++
		r1--

		l2, r2 := i, i+1
		for ; l2 >= 0 && r2 < len(s) && s[l2] == s[r2]; l2, r2 = l2-1, r2+1 {
		}
		l2++
		r2--

		if r1-l1 > right-left {
			left, right = l1, r1
		}
		if r2-l2 > right-left {
			left, right = l2, r2
		}
	}

	return s[left : right+1]
}

func main() {
	s := "bb"
	fmt.Println(longestPalindrome2(s))
}
