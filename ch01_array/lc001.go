package ch01_array

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, v := range nums {
		if j, ok := cache[target-v]; ok {
			return []int{j, i}
		}
		cache[v] = i
	}
	return []int{-1, -1}
}
