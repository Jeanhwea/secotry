package ch01_array

func rotate_01(a []int, k int) {
	n := len(a)
	b := make([]int, n)
	for i, v := range a {
		b[(i+k)%n] = v
	}
	copy(a, b)
}

func rotate_02(a []int, k int) {
	reverse := func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-i-1] = a[n-i-1], a[i]
		}
	}
	k = k % (len(a))
	reverse(a)
	reverse(a[:k])
	reverse(a[k:])
}

func rotate_03(a []int, k int) {
	// 辗转相除法求最大公约数
	gcd := func(x, y int) int {
		for x != 0 {
			x, y = y%x, x
		}
		return y
	}
	n := len(a)
	k %= n
	for first, count := 0, gcd(k, n); first < count; first++ {
		pre, curr := a[first], first
		for ok := true; ok; ok = curr != first {
			next := (curr + k) % n
			a[next], pre, curr = pre, a[next], next
		}
	}
}
