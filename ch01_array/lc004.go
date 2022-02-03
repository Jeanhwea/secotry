package ch01_array

func findMedianSortedArrays_01(a []int, b []int) float64 {
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
		return min(a[0], b[0])
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

func findKth_02(a, b []int, k int) int {
	i, j := 0, 0
	for {
		if i == len(a) {
			return b[j+k-1]
		}
		if j == len(b) {
			return a[i+k-1]
		}
		if k == 1 {
			return min(a[i], b[j])
		}
		i2 := min(i+k/2, len(a)) - 1
		j2 := min(j+k/2, len(b)) - 1
		if a[i2] <= b[j2] {
			k -= (i2 - i + 1)
			i = i2 + 1
		} else {
			k -= (j2 - j + 1)
			j = j2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

////////////////////////////////////////////////////////////////////////////////

func findMedianSortedArrays(a []int, b []int) float64 {
	n, m := len(a), len(b)
	if n > m {
		return findMedianSortedArrays(b, a)
	}
	k := (n + m) / 2

	for i+j+2 < k {
		gap := k - (i + j + 2)
		i += gap / 2
		if i >= n {
			i = n - 1
		}
		j = k - i - 2
	}
}
