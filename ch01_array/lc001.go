package ch01_array

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, val := range nums {
		if j, ok := cache[target-val]; ok {
			return []int{j, i}
		}
		cache[val] = i
	}
	return []int{-1, -1}
}
