package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var max, aLeft, bLeft int
	if len(a) > len(b) {
		max = len(a)
		bLeft = len(a) - len(b)
	} else {
		max = len(b)
		aLeft = len(b) - len(a)
	}

	result := ""
	var aVal, bVal, up int
	for i := max - 1; i >= 0; i-- {

		if i-aLeft < 0 {
			aVal = 0
		} else {
			aVal = int(a[i-aLeft]) - 48
		}
		if i-bLeft < 0 {
			bVal = 0
		} else {
			bVal = int(b[i-bLeft]) - 48
		}
		val := aVal + bVal + up
		if val == 3 {
			result = "1" + result
			up = 1
		} else if val == 2 {
			result = "0" + result
			up = 1
		} else {
			result = strconv.Itoa(val) + result
			up = 0
		}
	}

	if up == 1 {
		result = "1" + result
	}
	return result
}

func main() {
	a := "100"
	b := "110010"
	val := addBinary(a, b)
	fmt.Print(val)
}
