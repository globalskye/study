package main

import (
	"container/heap"
)

func main() {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
}

type Intheap struct {
	arr []int
}

func (h Intheap) Len() int {
	return len(h.arr)
}
func (h Intheap) Less(i, j int) bool {
	return h.arr[i] > h.arr[j]
}
func (h Intheap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Intheap) Push(x any) {
	h.arr = append(h.arr, x.(int))
}

func (h *Intheap) Pop() any {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &Intheap{arr: nums}

	heap.Init(h)
	res := 0

	for i := 0; i < k; i++ {

		res = heap.Pop(h).(int)

	}

	return res
}
