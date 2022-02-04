package ch01_array

import (
	"math"
	"sort"
)

func threeSumClosest(a []int, target int) int {
	sort.Ints(a)
	ans := math.MaxInt64
	for i := 0; i < len(a); i++ {
		k := len(a) - 1
		for j := i + 1; j < len(a); j++ {
			for j < k {
				sum := a[i] + a[j] + a[k]
				if sum > target {
					if abs(ans-target) > abs(sum-target) {
						ans = sum
					}
					k--
				} else {
					break
				}
			}
			if j < k && a[i]+a[j]+a[k] == target {
				return target
			}
			sum := a[i] + a[j] + a[k]
			if abs(ans-target) > abs(sum-target) {
				ans = sum
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
