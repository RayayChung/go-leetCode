package main

func divingBoard(shorter int, longer int, k int) []int {
	if shorter == longer {
		return []int{shorter * k}
	}

	interval := longer - shorter

	result := make([]int, k+1)
	result[0] = shorter * k

	for i := 1; i <= k; i++ {
		result[i] = result[i-1] + interval
	}
	return result
}
