package main

func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	var minArray, maxArray []int
	if len(A) > len(B) {
		minArray = B
		maxArray = A
	} else {
		minArray = A
		maxArray = B
	}

	maxMap := make(map[int][]int)
	for i := 0; i < len(maxArray); i++ {
		val := maxArray[i]
		maxMap[val] = append(maxMap[val], i)
	}

	var result int
	for j := 0; j < len(minArray); j++ {
		val := minArray[j]
		idxes := maxMap[val]

		if idxes == nil {
			continue
		}

		if result >= len(minArray)-j {
			break
		}
		for i := 0; i < len(idxes); i++ {
			idx := idxes[i]
			tmp := 0
			for k := j; k < len(minArray) && idx+k-j < len(maxArray); k++ {
				if minArray[k] == maxArray[idx+k-j] {
					tmp += 1
				} else {
					break
				}
			}
			if tmp > result {
				result = tmp
			}
		}
	}

	return result
}

//func main() {
//	a := []int{0,1,1,1,1}
//	b := []int{1,0,1,0,1}
//	val := findLength(a, b)
//	fmt.Print(val)
//}
