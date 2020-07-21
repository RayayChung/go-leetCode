package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) == 0 {
		if len(s1) == len(s2) && len(s1) == 0 {
			return true
		}
		return false
	}
	if len(s1) == 0 {
		return s3 == s2
	}
	if len(s2) == 0 {
		return s3 == s1
	}
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	result := make([][][]int, len(s3)+1)

	zero := make([][]int, 0)
	zero = append(zero, []int{0, 0})
	result[0] = zero

	for i := 1; i <= len(s3); i++ {

		val := s3[i-1]
		list := result[i-1]

		tmp := make([][]int, 0)
		for j := 0; j < len(list); j++ {
			subList := list[j]
			s1idx := subList[0]
			s2idx := subList[1]
			if subList[0] < len(s1) {
				first := s1[s1idx]
				if first == val {
					tmp = append(tmp, []int{s1idx + 1, subList[1]})
				}
			}
			if subList[1] < len(s2) {
				second := s2[s2idx]
				if second == val {
					tmp = append(tmp, []int{subList[0], s2idx + 1})
				}
			}
		}
		result[i] = tmp
	}

	l := result[len(s3)]

	if len(l) == 0 {
		return false
	}

	for k := 0; k < len(l); k++ {
		if l[k][0] == len(s1) && l[k][1] == len(s2) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
