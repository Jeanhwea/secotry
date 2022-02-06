package ch01_array

func trap(height []int) int {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0], rightMax[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func trap_03(height []int) int {
	i, j, leftMax, rightMax := 0, len(height)-1, 0, 0

	ans := 0
	for i < j {
		leftMax, rightMax = max(leftMax, height[i]), max(rightMax, height[j])
		if height[i] < height[j] {
			ans += leftMax - height[i]
			i++
		} else {
			ans += rightMax - height[j]
			j--
		}
	}

	return ans
}
