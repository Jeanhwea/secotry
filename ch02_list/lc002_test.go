package ch02_list

import "testing"

func Test_LC002_01(t *testing.T) {
	list1 := MakeList([]int{1, 2, 3})
	t.Logf("list1 = %v", ListToArr(list1))
}
