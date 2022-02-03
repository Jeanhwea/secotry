package ch01_array

import "testing"

func Test_LC189_01(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	res := []int{3, 4, 5, 6, 1, 2}
	rotate_03(nums, 4)
	for i := 0; i < len(nums); i++ {
		if nums[i] != res[i] {
			t.FailNow()
		}
	}
}
