package ch01_array

import (
	"sort"
)

// 100, 4, 200, 1, 3, 2
func longestConsecutive(a []int) int {
	cache := map[int]bool{}
	for _, v := range a {
		cache[v] = true
	}
	ans := 0
	for k := range cache {
		if !cache[k-1] { // 从起始数字开始往后查找
			count, curr := 1, k
			for cache[curr+1] {
				count++
				curr++
			}
			if count > ans {
				ans = count
			}
		}
	}
	return ans
}

func longestConsecutive_01(a []int) int {
	sort.Ints(a)
	ans, i := 0, 0
	for i < len(a) {
		count, j := 1, i+1
		for j < len(a) {
			if a[j] == a[j-1]+1 {
				count++
			}
			if a[j] > a[j-1]+1 {
				break
			}
			j++
		}
		if count > ans {
			ans = count
		}
		i = j
	}
	return ans
}
