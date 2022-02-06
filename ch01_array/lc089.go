package ch01_array

func grayCode(n int) []int {
	var ans []int
	total := 1 << n
	for i := 0; i < total; i++ {
		code := (i >> 1) ^ i
		ans = append(ans, code)
	}
	return ans
}
