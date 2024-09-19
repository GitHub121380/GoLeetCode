package main

import "fmt"

func main() {
	s := "abcaddefg"
	fmt.Println(lengthOfLongestSubstring(s))
}

/*
滑动窗口+一个可以记录窗口内字符是否出现过的容器，右移对应字符出现+1，左移对应字符出现-1
*/
func lengthOfLongestSubstring(s string) int {
	left, right, max := 0, 0, 0
	var freq [256]int
	for right < len(s) && left+max < len(s) {
		//这个0，其实是n-1，n表示子串中允许重复出现的次数，如果是1次，其实freq可以用bitset记录
		if freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}
