package ch01_array

func nextPermutation(nums []int) {
	// Step 1: 从右往左找到首次单调下降的位置，记做 i
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// Step 2: 在 nums[i+1:] 中，从右往左早到大于 nums[i] 的元素 nums[j], 如果找到交换
	if i >= 0 {
		j := len(nums) - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 3: 原地翻转 nums[i+1:]
	x, y := i+1, len(nums)-1
	for x < y {
		nums[x], nums[y] = nums[y], nums[x]
		x++
		y--
	}
}
