package main

func dailyTemperatures1(T []int) []int {

	result := make([]int, len(T))

	for i := 0; i < len(T)-1; i++ {
		length := 0
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				length = j - i
				break
			}
		}
		result[i] = length
	}

	return result
}

func dailyTemperatures2(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	var stack []int
	for i := 0; i < length; i++ {
		temperature := T[i]
		for len(stack) > 0 && temperature > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
