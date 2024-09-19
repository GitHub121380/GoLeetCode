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
// 本题变体1有求一个最接近0的三元组，那么和就要与0做查求绝对值，然后记录最小的那个三元组
// 本题变体2求指定target的，那么就计算头+尾与target-a的大小关系即可
// 本题变体3 变体1+变体2
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, len(nums))
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		//固定a后，双指针jk
		j := i + 1
		k := len(nums) - 1
		//开始单次查找和=-a的
		for j < k {
			//判重，起始不判
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k != len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			sum := nums[j] + nums[k]
			if sum > -nums[i] { //大了
				k--
			} else if sum < -nums[i] { //小了
				j++
			} else { //=
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return result
}
func main() {
	s := []int{-1, 0, 1, 2, -1, -4}
	threeSum(s)
}
