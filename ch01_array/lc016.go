package ch01_array

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return ans
}
