package ch01_array

import "testing"

func Test01_LC080(t *testing.T) {
	nums := []int{}
	len := removeDuplicates2(nums)
	if len != 0 {
		t.FailNow()
	}
}

func Test02_LC080(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	len := removeDuplicates2(nums)
	t.Logf("nums=%v, len=%v", nums, len)
	if len != 4 {
		t.FailNow()
	}
}

func Test03_LC080(t *testing.T) {
	nums := []int{1, 2, 2, 2, 3}
	len := removeDuplicates2(nums)
	t.Logf("nums=%v, len=%v", nums, len)
	if len != 4 {
		t.FailNow()
	}
}

func Test04_LC080(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	len := removeDuplicates2(nums)
	t.Logf("nums=%v, len=%v", nums, len)
	if len != 5 {
		t.FailNow()
	}
}
