package ch01_array

import "testing"

func Test_LC018_01(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	ans := fourSum(nums, 0)
	t.Logf("ans = %v", ans)
}

func Test_LC018_02(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	ans := fourSum(nums, 8)
	t.Logf("ans = %v", ans)
}
