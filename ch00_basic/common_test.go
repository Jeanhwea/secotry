package ch00_basic

import (
	"sort"
	"testing"
)

func TestCommon01(t *testing.T) {
	a := []int{3, 5, 2, 0, 1, 1, 9}
	sort.Ints(a)
	t.Logf("a = %v", a)

	b := []float64{3, 1.0, 0, 9.4}
	sort.Float64s(b)
	t.Logf("b = %v", b)
}
