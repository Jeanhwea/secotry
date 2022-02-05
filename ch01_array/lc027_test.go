package ch01_array

import "testing"

func Test_LC027_01(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	len := removeElement(nums, 3)
	t.Logf("nums = %v, len = %v", nums, len)
}
