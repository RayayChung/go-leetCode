package main

import (
	"sort"
)

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	result := make([][]int, length)

	for i := 0; i < length; i++ {
		result[i] = make([]int, 0)
	}

	for i := 0; i < length; i++ {
		if i == 0 {
			result[0] = append(result[0], triangle[0][0])
			continue
		}
		for j := 0; j < len(triangle[i]); j++ {
			row := triangle[i]
			var val int
			if j < len(result[i-1]) {
				val = result[i-1][j]
				if j-1 >= 0 {
					val = minValue(val, result[i-1][j-1])
				}
			} else {
				val = result[i-1][j-1]
			}
			result[i] = append(result[i], val+row[j])
		}
	}
	lastRow := triangle[length-1]
	sort.Ints(lastRow)
	return lastRow[0]
}

func minValue(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

//func main() {
//	val := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
//	minimumTotal(val)
//}
