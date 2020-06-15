package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		if i >= 1 && current == nums[i-1] {
			continue
		}

		last := len(nums) - 1

		for first := i + 1; first < len(nums); first++ {

			if first > i+1 && nums[first] == nums[first-1] {
				continue
			}

			for first < last && nums[first]+nums[last] > 0 {
				last--
			}

			if first == last {
				continue
			}

			if nums[first]+nums[last] == 0 {
				result = append(result, []int{current, nums[first], nums[last]})
			}
		}
	}
	return result
}
