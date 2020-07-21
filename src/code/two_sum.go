package main

func twoSum(numbers []int, target int) []int {

	length := len(numbers)

	if length <= 1 {
		return []int{}
	}
	i := 0
	j := length - 1

	for i < j {
		val := numbers[i] + numbers[j]
		if val == target {
			return []int{i + 1, j + 1}
		} else if val > target {
			j--
		} else if val < target {
			i++
		}
	}
	return []int{}
}
