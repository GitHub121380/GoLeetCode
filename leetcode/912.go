package main

import (
	"fmt"
	"math/rand"
)

// 自己实现排序
func sortArray(nums []int) []int {
	return quickSort(nums)
}

func quickSort(nums []int) []int {
	partation(nums, 0, len(nums)-1)
	return nums
}

func partation(nums []int, l, r int) {
	if l >= r {
		return
	}

	index := l + rand.Intn(r-l+1)
	nums[l], nums[index] = nums[index], nums[l]

	i := l
	j := r
	//key := nums[l]

	for i < j {
		for i < j && nums[j] >= nums[l] {
			j--
		}

		for i < j && nums[i] <= nums[l] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	//i==j 交点
	nums[l], nums[j] = nums[j], nums[l]
	partation(nums, l, j-1)
	partation(nums, j+1, r)
}
func main() {
	s := []int{
		4, 5, 2, 3, 1,
	}
	fmt.Println(quickSort(s))
}
