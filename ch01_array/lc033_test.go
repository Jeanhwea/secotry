package ch01_array

import "testing"

func Test01_LC033(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	k := search(nums, 0)
	if k != 4 {
		t.FailNow()
	}
}
