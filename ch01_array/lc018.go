package ch01_array

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	if len(nums) < 4 {
		return ans
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j-1] == nums[j] {
				continue
			}

			x, y := j+1, len(nums)-1
			for x < y {
				if x != j+1 && nums[x-1] == nums[x] {
					x++
					continue
				}
				sum := nums[i] + nums[j] + nums[x] + nums[y]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[x], nums[y]})
					x++
				} else if sum > target {
					y--
				} else {
					x++
				}
			}
		}
	}
	return ans
}
