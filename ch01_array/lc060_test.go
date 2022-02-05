package ch01_array

import "testing"

func Test_LC060_01(t *testing.T) {
	ans := getPermutation(4, 13)
	t.Logf("ans = %v", ans)
}

func Test_LC060_02(t *testing.T) {
	ans := getPermutation(4, 24)
	t.Logf("ans = %v", ans)
}
