package ch00_basic

import "testing"

func TestArray01(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	i := 2
	t.Logf("a[:i]=%v, a[i:]=%v", a[:i], a[i:])
}
