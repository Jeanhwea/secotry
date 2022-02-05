package ch01_array

import "strconv"

// Cantor Expansion
// k-1 = an*(n-1)! + ... + a2*2! + a1
func getPermutation(n int, k int) string {
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = i * fact[i-1]
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i <= n; i++ {
		valid[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		order := k/fact[i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= fact[i]
	}

	return ans
}
