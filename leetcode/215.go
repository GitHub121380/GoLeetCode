package main

import "math/rand"

// 位图法
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	freq := [20001]int{0}

	for i := 0; i < len(nums); i++ {
		freq[nums[i]+10000]++
	}

	for i := len(freq) - 1; i >= 0; i-- {
		if freq[i] >= k {
			return i - 10000
		} else {
			k = k - freq[i]
		}
	}
	return 0
}

// 快排中的分治，只需要让右边的数都大于自己，然后确定到k即可
// https://blog.csdn.net/SUICIDEKILL/article/details/137060375
func findKthLargest2(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}

	//均衡操作，取随机数做partition
	randIndex := l + rand.Intn(r-l+1)
	nums[l], nums[randIndex] = nums[randIndex], nums[l]
	partition := nums[l]

	i := l
	j := r
	for i < j {
		for i < j && nums[j] >= partition {
			j--
		}
		for i < j && nums[i] <= partition {
			i++
		}

		//上面都把大的和小的都过完了，然后再互换 然后再走下一趟
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	//i和j聚在一起了 此时说明i==j，交点处是区间中位数，需要将partition换过来
	nums[l], nums[j] = nums[j], nums[l]

	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

// 快排的非递归
func findKthLargest3(nums []int, k int) int {
	n := len(nums)
	return quickselect2(nums, 0, n-1, n-k)
}

func quickselect2(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}

	//均衡操作，取随机数做partition
	randIndex := l + rand.Intn(r-l+1)
	nums[l], nums[randIndex] = nums[randIndex], nums[l]
	partition := nums[l]

	i := l
	j := r
	for i < j {
		for i < j && nums[j] >= partition {
			j--
		}
		for i < j && nums[i] <= partition {
			i++
		}

		//上面都把大的和小的都过完了，然后再互换 然后再走下一趟
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	//i和j聚在一起了 此时说明i==j，交点处是区间中位数，需要将partition换过来
	nums[l], nums[j] = nums[j], nums[l]

	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}
