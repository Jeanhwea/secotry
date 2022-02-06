package ch01_array

import "testing"

func Test_LC042_01(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans := trap(height)
	// t.Logf("ans = %v", ans)
	if ans != 6 {
		t.FailNow()
	}
}
