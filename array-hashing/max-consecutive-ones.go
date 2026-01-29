package arrayhashing

import "sort"

func FindMaxConsecutiveOnes(nums []int) int {
	var count int
	var result []int

	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			result = append(result, count)
			count = 0
		}
	}

	result = append(result, count)
	sort.Ints(result)
	return result[len(result)-1]
}
