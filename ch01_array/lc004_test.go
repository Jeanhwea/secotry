package ch01_array

import "testing"

func Test_LC004_01(t *testing.T) {
	num1 := []int{1, 2, 3, 4, 5}
	num2 := []int{3, 4, 4, 5}
	val := findKth(num1, num2, 2)
	if val != 2 {
		t.FailNow()
	}
}

func Test_LC004_02(t *testing.T) {
	num1 := []int{1, 2}
	num2 := []int{3, 4}
	val := findKth(num1, num2, 2)
	if val != 2 {
		t.FailNow()
	}
	val2 := findKth(num1, num2, 3)
	if val2 != 3 {
		t.FailNow()
	}
}

func Test_LC004_03(t *testing.T) {
	num1 := []int{1, 3}
	num2 := []int{2}
	val := findMedianSortedArrays(num1, num2)
	if val != 2 {
		t.FailNow()
	}
}
