package arrayhashing

func TwoSum(nums []int, target int) []int {
	var sum int
	var result []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == target {
				result = append(result, i, j)
				break
			}
		}
	}
	return result
}
