package ch01_array

func search(a []int, target int) int {
	beg, end := 0, len(a)-1

	// [beg, end]
	for beg <= end {
		mid := beg + (end-beg)/2
		if a[mid] == target {
			return mid
		}
		if a[mid] < a[beg] {
			if a[mid] < target && target <= a[end] {
				beg = mid + 1
			} else {
				end = mid - 1
			}
		} else { // a[mid] > a[beg], 因为 a[mid] 不可能等于 a[beg]
			if a[beg] <= target && target < a[mid] {
				end = mid - 1
			} else {
				beg = mid + 1
			}
		}
	}

	return -1
}
