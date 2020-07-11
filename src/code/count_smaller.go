package main

func countSmaller(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		var val int
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				val++
			}
		}
		result[i] = val
	}

	return result
}
