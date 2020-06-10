package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	numStr := fmt.Sprintf("%d", x)

	length := len(numStr)
	for i := 0; i < length; i++ {
		if numStr[i:i+1] != numStr[length-1-i:length-i] {
			return false
		}
	}
	return true
}
