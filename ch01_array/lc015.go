package ch01_array

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j-1] == nums[j] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			if j < k && nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}

func threeSum_02(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		// 使用双指针来查找
		j, k := i+1, len(nums)-1
		for j < k {
			if j != i+1 && nums[j-1] == nums[j] {
				j++
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
			if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return ans
}
