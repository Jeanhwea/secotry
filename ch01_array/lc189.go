package ch01_array

func rotate_01(a []int, k int) {
	n := len(a)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[(i+k)%n] = a[i]
	}

	for i := 0; i < n; i++ {
		a[i] = b[i]
	}
}

func rotate_02(a []int, k int) {
	n, k := len(a), k%(len(a))
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
	for i := 0; i < k/2; i++ {
		a[i], a[k-i-1] = a[k-i-1], a[i]
	}
	for i := 0; i < (n-k)/2; i++ {
		a[k+i], a[n-i-1] = a[n-i-1], a[k+i]
	}
}
