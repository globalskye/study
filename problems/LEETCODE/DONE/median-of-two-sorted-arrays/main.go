package main

import "fmt"

func main() {
	findMedianSortedArrays([]int{1, 2, 3}, []int{1, 2, 4, 4, 5})
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tmp := make([]int, len(nums1)+len(nums2))
	idx := 0
	i := 0
	j := 0
	for idx < len(tmp) {

		if j == len(nums2) {
			tmp[idx] = nums1[i]
			idx++
			i++
			continue
		}
		if i == len(nums1) {
			tmp[idx] = nums2[j]
			idx++
			j++
			continue
		}
		if nums1[i] < nums2[j] {
			tmp[idx] = nums1[i]
			idx++
			i++
		} else {
			tmp[idx] = nums2[j]
			idx++
			j++
		}

	}
	fmt.Println(tmp)
	if len(tmp)%2 == 0 {
		a := float64(tmp[len(tmp)/2])
		b := float64(tmp[len(tmp)/2-1])

		return (a + b) / 2

	}
	return float64(tmp[len(tmp)/2])

}
