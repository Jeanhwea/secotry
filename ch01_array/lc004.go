package ch01_array

import "math"

func findMedianSortedArrays_01(a []int, b []int) float64 {
	total := len(a) + len(b)
	if total%2 == 0 {
		return float64(findKth(a, b, total/2)+findKth(a, b, total/2+1)) / 2
	}
	return float64(findKth(a, b, total/2+1))
}

// 递归版
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

// 非递归版本
func findKth_02(a, b []int, k int) int {
	i, j, n, m := 0, 0, len(a), len(b)
	for {
		if i == n {
			return b[j+k-1]
		}
		if j == m {
			return a[i+k-1]
		}
		if k == 1 {
			return min(a[i], b[j])
		}
		x, y := min(i+k/2, n)-1, min(j+k/2, m)-1
		if a[x] < b[y] {
			k -= x - i + 1
			i = x + 1
		} else {
			k -= y - j + 1
			j = y + 1
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

func findMedianSortedArrays(a []int, b []int) float64 {
	n, m := len(a), len(b)
	if n > m {
		return findMedianSortedArrays(b, a)
	}

	beg, end, leftMaxVal, rightMinVal := 0, n, 0, 0
	for beg <= end {
		// i+j == (m+n+1)/2 使得 a[0..i] + b[0..j] 刚好是 a[..] 并 b[..] 总元素的一半
		i := (beg + end) / 2
		j := (m+n+1)/2 - i

		// 记录区间边界值 x0, x1, y0, y1 => a[i-1], a[i], b[j-1], b[j]
		x0, x1, y0, y1 := math.MinInt64, math.MaxInt64, math.MinInt64, math.MaxInt64
		if i != 0 {
			x0 = a[i-1]
		}
		if i != n {
			x1 = a[i]
		}
		if j != 0 {
			y0 = b[j-1]
		}
		if j != m {
			y1 = b[j]
		}

		// leftMaxVal 是左半部分的最大值， rightMinVal 是右半部分的最小值
		if x0 <= y1 {
			leftMaxVal, rightMinVal = max(x0, y0), min(x1, y1)
			beg = i + 1
		} else {
			end = i - 1
		}
	}

	if (n+m)%2 == 0 {
		return float64(leftMaxVal+rightMinVal) / 2
	}
	return float64(leftMaxVal)
}
