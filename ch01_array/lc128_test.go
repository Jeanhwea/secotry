package ch01_array

import "testing"

func Test_LC128_01(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	ans := longestConsecutive(nums)
	t.Logf("ans = %v", ans)
}

func Test_LC128_02(t *testing.T) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	ans := longestConsecutive(nums)
	t.Logf("ans = %v", ans)
}
