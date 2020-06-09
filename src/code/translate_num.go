package main

import (
	"fmt"
)

func translateNum(num int) int {
	numStr := fmt.Sprintf("%d", num)

	f1 := 0
	f2 := 0
	result := 1

	for i := 0; i < len(numStr); i++ {
		f1 = f2
		f2 = result
		result = f2

		if i == 0 {
			continue
		}

		val := numStr[i-1 : i+1]

		fmt.Println(val)

		if val <= "25" && val >= "10" {
			result += f1
		}
	}
	return result
}
