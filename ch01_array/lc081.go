package ch01_array

func search2(a []int, target int) bool {
	beg, end := 0, len(a)-1

	// [beg, end]
	for beg <= end {
		mid := beg + (end-beg)/2
		if a[mid] == target {
			return true
		}
		if a[mid] < a[beg] {
			if a[mid] < target && target <= a[end] {
				beg = mid + 1
			} else {
				end = mid
			}
		} else if a[mid] > a[beg] {
			if a[beg] <= target && target < a[mid] {
				end = mid - 1
			} else {
				beg = mid
			}
		} else { // 等值时可能死循环，需要特殊处理
			beg++
		}
	}

	return false
}
