package ch01_array

import "testing"

func Test_LC015_01(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	t.Logf("ans = %v", ans)
}
