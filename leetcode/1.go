package main

// 使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N) 降低到 O(1)。
func twoSum(nums []int, target int) []int {
	//key是nums值， value是下标，因为我们要返回下标
	m := make(map[int]int, len(nums))
	//我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，
	//然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。
	for i := 0; i < len(nums); i++ {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{i, idx}
		}
		m[nums[i]] = i
	}
	return nil
}
