package ch01_array

func plusOne(digits []int) []int {
	ans := []int{}
	carry, digit := 1, 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		carry, digit = sum/10, sum%10
		ans = append([]int{digit}, ans...)
	}

	if carry > 0 {
		ans = append([]int{carry}, ans...)
	}

	return ans
}
