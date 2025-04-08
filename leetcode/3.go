package main

import "fmt"

func main() {
	s := "abcaddefg"
	fmt.Println(lengthOfLongestSubstring(s))
}

/*
滑动窗口+一个可以记录窗口内字符是否出现过（出现过几次）的容器，没出现过就r右移，次数+1；出现过就l右移，次数-1
*/
func lengthOfLongestSubstring(s string) int {
	var l, r, max int

	//需要一个存储当前滑动窗口中字符出现次数的容器
	/*set 也行 map也行 直接用一个arr也行（下标表示字符，值表示出现次数）*/
	var freq [512]int
	for r < len(s) && l+max < len(s) {
		//这个0，其实是n-1，n表示子串中允许重复出现的次数，这样就是各种变种题
		if freq[s[r]] == 0 { //r在当前滑动窗口没出现过，则右移r，然后现在的s[r]是下次判断的
			freq[s[r]]++
			r++
		} else { //r在当前滑动窗口出现过，则l左移，抛弃l尝试新的子串，直到r在当前滑动窗口没出现过
			freq[s[l]]--
			l++
		}
		//这里不加1是因为上面r++了，现在的r其实是当时读取r的+1；而如果是l++，那么肯定不会比max还大
		if r-l > max {
			max = r - l
		}
	}
	return max
}
