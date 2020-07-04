package main

import "sort"

type KMinHeap []int

func (h KMinHeap) Len() int {
	return len(h)
}

func (h KMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h KMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *KMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *KMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	row, column := len(matrix), len(matrix[0])

	result := make([]int, row*column)

	index := 0
	for _, row := range matrix {
		for _, column := range row {
			result[index] = row[column]
			index++
		}
	}

	sort.Ints(result)

	return result[k-1]
}
