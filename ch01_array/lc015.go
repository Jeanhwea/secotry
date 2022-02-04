package ch01_array

import (
	"sort"
)

func threeSum(a []int) [][]int {
	sort.Ints(a)
	var ans [][]int
	for i := 0; i < len(a); i++ {
		if i != 0 && a[i-1] == a[i] {
			continue
		}
		k := len(a) - 1
		for j := i + 1; j < len(a); j++ {
			if j != i+1 && a[j-1] == a[j] {
				continue
			}
			for j < k && a[i]+a[j]+a[k] > 0 {
				k--
			}
			if j < k && a[i]+a[j]+a[k] == 0 {
				ans = append(ans, []int{a[i], a[j], a[k]})
			}
		}
	}
	return ans
}
