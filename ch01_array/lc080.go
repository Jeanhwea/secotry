package ch01_array

func removeDuplicates2(a []int) int {
	k := -1
	for i := 0; i < len(a); i++ {
		if k > 0 && a[i] == a[k] && a[i] == a[k-1] {
			continue
		}
		k++
		a[k] = a[i]
	}
	return k + 1
}
