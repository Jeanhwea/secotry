package ch01_array

import "testing"

func Test01_LC026(t *testing.T) {
	nums := []int{}
	len := removeDuplicates(nums)
	if len != 0 {
		t.FailNow()
	}
}

func Test02_LC026(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	len := removeDuplicates(nums)
	t.Logf("nums=%v, len=%v", nums, len)
	if len != 3 {
		t.FailNow()
	}
}
