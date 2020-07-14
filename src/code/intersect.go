package main

func intersect(nums1 []int, nums2 []int) []int {
	//sort.Ints(nums1)
	//sort.Ints(nums2)
	//
	//var i, j int
	//
	//result := make([]int, 0)
	//for i < len(nums1) && j < len(nums2) {
	//	if nums1[i] < nums2[j] {
	//		i++
	//	} else if nums1[i] > nums2[j] {
	//		j++
	//	} else {
	//		result = append(result, nums2[j])
	//		i++
	//		j++
	//	}
	//}
	//
	//return result

	countMap := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		countMap[nums1[i]] += 1
	}

	result := make([]int, 0)

	for j := 0; j < len(nums2); j++ {
		count := countMap[nums2[j]]

		if count > 0 {
			result = append(result, nums2[j])
			countMap[nums2[j]] = count - 1
		} else {
			continue
		}
	}

	return result
}
