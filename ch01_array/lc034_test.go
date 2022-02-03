package ch01_array

import "testing"

func Test_LC034_01(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	ans := searchRange(nums, 8)
	t.Logf("ans = %v", ans)
}

func Test_LC034_02(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	ans := searchRange(nums, 6)
	t.Logf("ans = %v", ans)
}
