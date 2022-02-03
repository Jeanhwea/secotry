package ch01_array

import "testing"

func Test_LC080_01(t *testing.T) {
	nums := []int{}
	len := removeDuplicates2(nums)
	if len != 0 {
		t.FailNow()
	}
}

func Test_LC080_02(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	len := removeDuplicates2(nums)
	t.Logf("nums=%v, len=%v", nums, len)
	if len != 4 {
		t.FailNow()
	}
}

func Test_LC080_03(t *testing.T) {
	nums := []int{1, 2, 2, 2, 3}
	len := removeDuplicates2(nums)
	t.Logf("nums=%v, len=%v", nums, len)
	if len != 4 {
		t.FailNow()
	}
}

func Test_LC080_04(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	len := removeDuplicates2(nums)
	t.Logf("nums=%v, len=%v", nums, len)
	if len != 5 {
		t.FailNow()
	}
}
