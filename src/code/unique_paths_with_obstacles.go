package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	row := len(obstacleGrid)
	column := len(obstacleGrid[0])
	result := make([][]int, row)

	for i := 0; i < row; i++ {
		result[i] = make([]int, column)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == 0 && j == 0 {
				result[i][j] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
			} else {
				if i-1 >= 0 {
					result[i][j] += result[i-1][j]
				}
				if j-1 >= 0 {
					result[i][j] += result[i][j-1]
				}
			}
		}
	}

	return result[row-1][column-1]
}

func main() {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	val := uniquePathsWithObstacles(input)
	fmt.Print(val)
}
