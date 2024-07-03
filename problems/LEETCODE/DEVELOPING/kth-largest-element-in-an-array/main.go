package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
func findKthLargest(nums []int, k int) int {

	intMax := 0
	for i := 0; i < len(nums); i++ {
		heap.Init()
	}

	return intMax
}
