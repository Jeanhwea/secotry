package ch01_array

import "sort"

func searchRange(a []int, target int) []int {
	beg, end := 0, len(a)-1
	for beg <= end {
		mid := beg + (end-beg)/2
		if a[mid] == target {
			leftMost, rightMost := mid, mid
			for leftMost > beg && a[leftMost-1] == target {
				leftMost--
			}
			for rightMost < end && a[rightMost+1] == target {
				rightMost++
			}
			return []int{leftMost, rightMost}
		}
		if a[mid] < target {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return []int{-1, -1}
}

func searchRange_01(a []int, target int) []int {
	leftMost := sort.SearchInts(a, target)
	if leftMost == len(a) || a[leftMost] != target {
		return []int{-1, -1}
	}
	rightMost := sort.SearchInts(a, target+1) - 1
	return []int{leftMost, rightMost}
}
