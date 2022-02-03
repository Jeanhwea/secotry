package ch01_array

func findMedianSortedArrays(a []int, b []int) float64 {
	total := len(a) + len(b)
	if total%2 == 0 {
		return float64(findKth(a, b, total/2)+findKth(a, b, total/2+1)) / 2
	}
	return float64(findKth(a, b, total/2+1))
}

func findKth(a, b []int, k int) int {
	n, m := len(a), len(b)
	if n > m {
		return findKth(b, a, k)
	}

	if n <= 0 {
		return b[k-1]
	}
	if k <= 1 {
		if b[0] < a[0] {
			return b[0]
		} else {
			return a[0]
		}
	}

	i := k/2 - 1
	if i >= n {
		i = n - 1
	}
	j := k - i - 2
	if a[i] < b[j] {
		return findKth(a[i+1:], b, k-i-1)
	} else {
		return findKth(a, b[j+1:], k-j-1)
	}
}
