package main

import (
	"sort"
)

func countSmaller(nums []int) []int {
	//result := make([]int, len(nums))
	//
	//for i := 0; i < len(nums); i++ {
	//	var val int
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[j] < nums[i] {
	//			val++
	//		}
	//	}
	//	result[i] = val
	//}
	//
	//return result

	dMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		dMap[nums[i]] = 1
	}

	dArray := make([]int, len(dMap))
	result := make([]int, len(nums))
	count := make([]int, len(nums))
	var idx int
	for key := range dMap {
		dArray[idx] = key
		idx++
	}

	sort.Ints(dArray)

	for i := range dArray {
		dMap[dArray[i]] = i
	}

	for j := len(nums) - 1; j >= 0; j-- {
		index := dMap[nums[j]]
		var val int
		for i := 0; i < index; i++ {
			val += count[i]
		}
		count[index] += 1
		result[j] = val
	}

	return result
}

//func main() {
//	for i := 0; i < 10; i++ {
//		fmt.Println(i & -i)
//	}
//}
