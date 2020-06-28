package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= s {
			return 1
		} else {
			return 0
		}
	}

	min := len(nums) + 1
	head := 0
	next := 1
	sum := nums[0] + nums[1]
	fmt.Printf("sum: %d\n", sum)
	for head < next && head < len(nums) && next < len(nums) {
		if nums[head] >= s || nums[next] >= s {
			return 1
		}

		if sum >= s {
			if next-head+1 < min {
				min = next - head + 1
				fmt.Printf("%d-%d-%d\n", head, next, sum)
			}
			sum -= nums[head]
			head++
		} else {
			next++
			if next == len(nums) {
				break
			}
			sum += nums[next]
		}
		fmt.Printf("sum: %d\n", sum)
	}

	if min == len(nums)+1 {
		return 0
	}
	return min
}

func main() {
	arrayLen := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
	fmt.Println(arrayLen)
}
