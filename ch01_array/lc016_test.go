package ch01_array

import "testing"

func Test_LC016_01(t *testing.T) {
	nums := []int{-1, 2, 1, 4}
	ans := threeSumClosest(nums, 1)
	t.Logf("ans = %v", ans)
}
