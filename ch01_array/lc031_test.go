package ch01_array

import "testing"

func Test_LC031_01(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Logf("prev = %v", nums)
	nextPermutation(nums)
	t.Logf("next = %v", nums)
}

func Test_LC031_02(t *testing.T) {
	nums := []int{3, 2, 1}
	t.Logf("prev = %v", nums)
	nextPermutation(nums)
	t.Logf("next = %v", nums)
}

func Test_LC031_03(t *testing.T) {
	nums := []int{1, 5, 1}
	t.Logf("prev = %v", nums)
	nextPermutation(nums)
	t.Logf("next = %v", nums)
}
