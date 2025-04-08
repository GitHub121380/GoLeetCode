package main

import "sort"

//排序+双指针
//排序后从左到右以此增大

// 固定一个元素a，然后看a右边的元素。我们要从右面的里面找两个相加等于-a（为什么不找左边的，因为左边的已经在之前走过同样的逻辑了，0号元素除外）
// 这些元素定义一个头指针一个尾指针，尾肯定大于等于头的因为排过序了。
// 如果头+尾>-a，说明和大了需要变小，尾左移动减小；如果头+尾<-a，说明和小了需要增加，头向右移动增大；头+尾=-a，说明满足条件，添加三元组并同时头右移尾左移；直到头=尾，break；
// 头和头-1不能相等，尾和尾+1也不能相等，否则继续移动
// 头尾这样操作一次是n，然后a每次这样执行完一遍，a指针向右移动直到倒数第二个break
// 共计n^2时间复杂度
// 本题变体1有求一个最接近0的三元组，那么和就要与0做差求绝对值，然后记录最小的那个三元组
// 本题变体2求指定target的，那么就计算头+尾与target-a的大小关系即可
// 本题变体3 变体1+变体2
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, len(nums))
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, l-1
		for left < right {

			if left > i+1 && nums[left] == nums[left-1] { //左侧重复了
				left++
				continue
			}
			if right < l-1 && nums[right] == nums[right+1] { //右侧重复了
				right--
				continue
			}

			if nums[left]+nums[right] == -nums[i] {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if nums[left]+nums[right] > -nums[i] {
				right--
			} else if nums[left]+nums[right] < -nums[i] {
				left++
			}
		}

	}
	return result
}

// 其实我感觉可以用leet1两数之和的解法，因为三数之和就是两数之和的变体
// 将每个元素的负数和下标存下来，然后排序，再用头尾指针求和看相等并且下标不能相等，这样时间复杂度也就是n2
// 但是这样有一个多余的空间消耗，而且还需要赋值读取
func test15_1(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, len(nums))

	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[-nums[i]] = i
	}

	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if idx, ok := m[(nums[i] + nums[j])]; ok && j < idx {
				result = append(result, []int{nums[i], nums[j], nums[idx]})
			}
		}
	}
	return result
}
func main() {
	//s := []int{-1, 0, 1, 2, -1, -4}
	s := []int{0, 0, 0, 0}
	//threeSum(s)
	test15_1(s)
}
