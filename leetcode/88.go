package main

import "fmt"

/*
88. 合并两个有序数组
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/

// 最简单是直接全放到num1上，然后调用sort函数，时间复杂度o(m+n)log(m+n)
// 第二简单就是再创建一个新arr，然后两个从头往后遍历，最后再把新arr放到num1上
// 但是如果要求空间复杂度o1的话，就只能直接在num1上操作了
// 可以从大到小尾指针遍历，然后从尾巴赋值nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		//nums1 = nums2
		return
	}

	i, j := m-1, n-1
	p := m + n - 1

	for p >= 0 {
		if j < 0 {
			return
		}
		if i < 0 {
			nums1[p] = nums2[j]
			p--
			j--
			continue
		}
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			nums1[i] = -(1 << 10)
			i--
			p--
		} else {
			nums1[p] = nums2[j]
			nums2[j] = -(1 << 10)
			j--
			p--
		}
	}
}

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}
