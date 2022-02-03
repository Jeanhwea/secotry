package ch00_basic

import (
	"sort"
	"testing"
)

func TestSort01(t *testing.T) {
	// 排序整数
	a := []int{3, 5, 2, 0, 1, 1, 9}
	sort.Ints(a)
	t.Logf("a = %v", a)

	// 排序浮点数
	b := []float64{3, 1.0, 0, 9.4}
	sort.Float64s(b)
	t.Logf("b = %v", b)

	// 排序字符串
	s := []string{"apple", "pear", "banana"}
	sort.Strings(s)
	t.Logf("s = %v", s)
}

func TestSort02(t *testing.T) {
	a := []int{3, 5, 2, 0, 1, 1, 9}
	// ans1 := sort.SearchInts(a, 5)
	// t.Logf("ans1 = %v", ans1)
	sort.Ints(a)
	ans2 := sort.SearchInts(a, 1) // a 必须是排序好的
	t.Logf("ans2 = %v", ans2)
}
