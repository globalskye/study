package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}
func merge(nums1 []int, m int, nums2 []int, n int) []int {

	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	return nums1
}
