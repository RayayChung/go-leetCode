package main

func maxScoreSightseeingPair(A []int) int {
	max := A[0]
	var ans int
	for i := 1; i < len(A); i++ {
		max = maxValue(max, A[i-1]+i-1)
		ans = max + A[i] - i
	}
	return ans
}

func maxValue(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
