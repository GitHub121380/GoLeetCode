package main

// 时间复杂度为 O(log n)
func search(nums []int, target int) int {
	//	找到对应的k，可是这样最大时间复杂度有n，然后再用二分查找是logn  但是这一步我n我们就可以直接找到了
	k := 0
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}
		if nums[i] > nums[i+1] {
			k = i
			break
		}
	}

	//	k的左边和右边分别有序  二分查找是logn
	//左边
	l := 0
	r := k
	mid := -1
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	//右边
	l = k + 1
	r = len(nums) - 1
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 直接二分查找
// 由于数组基本有序，虽然中间有⼀个“断开点”，还是可以使⽤⼆分搜索的算法来实现。
// 现在数组前⾯⼀段是数值⽐较⼤的数，后⾯⼀段是数值偏⼩的数。
// 如果 mid 落在了前⼀段数值⽐较⼤的被旋转的区间内了，那么mid在这个区间内升序，那么⼀定有 nums[mid] > nums[low]
// 如果mid落在后⾯⼀段数值⽐较⼩的没有旋转区间内，那么前面被旋转的区间包括最左侧的肯定都比它大，那么⼀定有nums[mid] ≤ nums[low] 。
// 如果 nums[mid] < nums[high] ，如果是落在前⾯⼀段
// 数值⽐较⼤的区间内，
// nums[mid] ≤ nums[high] 。还有
// nums[low] == nums[mid] 和
// nums[mid] 的情况，单独处理即可。最后找到则输出 mid，没有找到则输出 -1
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1
	mid := -1
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] { //mid在被旋转的大区间中
			//后面那个条件如果不加的话，只能说明target比nums[mid]小，但是无法确定是在升序大区间中mid左边，还是在小区间中mid右边，无法移动左或右指针
			//加上就表示target在l~mid中间 否则就只可能是在小区间中了
			if nums[mid] > target && nums[l] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//在右侧有序区间
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
