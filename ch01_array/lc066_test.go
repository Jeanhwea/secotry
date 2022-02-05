package ch01_array

import "testing"

func Test_LC066_01(t *testing.T) {
	digits := []int{1, 2, 3}
	ans := plusOne(digits)
	t.Logf("ans = %v", ans)
}

func Test_LC066_02(t *testing.T) {
	digits := []int{9, 9}
	ans := plusOne(digits)
	t.Logf("ans = %v", ans)
}
