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

func TestArray03(t *testing.T) {
	arr := []string{"s1", "s2", "s3", "s4", "s5"}
	remove := []string{"s1", "s4"}
loop:
	for i := 0; i < len(arr); i++ {
		url := arr[i]
		for _, rem := range remove {
			if url == rem {
				arr = append(arr[:i], arr[i+1:]...)
				i--
				continue loop
			}
		}
	}
	t.Logf("arr = %v", arr)
}
