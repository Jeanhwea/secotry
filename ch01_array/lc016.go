package ch01_array

import (
	"math"
	"sort"
)

func threeSumClosest(a []int, target int) int {
	sort.Ints(a)
	ans := math.MaxInt32

	for i := 0; i < len(a); i++ {
		j, k := i+1, len(a)-1
		for j < k {
			sum := a[i] + a[j] + a[k]
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
