package ch00_basic

import "testing"

func TestArray01(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	i := 2
	t.Logf("a[:i]=%v, a[i:]=%v", a[:i], a[i:])
}

func TestArray02(t *testing.T) {
	a := []int{1, 2}
	// slice 前插法
	a = append([]int{0}, a...)
	t.Logf("a = %v", a)

	// slice 后插法
	a = append(a, 3)
	t.Logf("a = %v", a)
}
