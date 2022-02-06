package ch01_array

import (
	"math"
)

func trap(height []int) int {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0], rightMax[n-1] = math.MinInt32, math.MinInt32
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}

	ans := 0
	for i := 0; i < n; i++ {
		curr := min(leftMax[i], rightMax[i]) - height[i]
		ans += max(0, curr)
	}
	return ans
}
