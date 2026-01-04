package twopointers

import "math"

func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			result[i] = nums[right] * nums[right]
			right--
		} else {
			result[i] = nums[left] * nums[left]
			left++
		}
	}

	return result
}
