package main

import (
	"fmt"
)

func main() {

	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 6))
}
func search(nums []int, target int) int {

	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer <= rightPointer {
		middlePointer := (leftPointer + rightPointer) / 2
		currentValue := nums[middlePointer]

		if currentValue == target {
			return middlePointer
		}

		if nums[middlePointer] < nums[rightPointer] && nums[middlePointer] > nums[leftPointer] {

		}
		if nums[middlePointer] > nums[leftPointer] {
			leftPointer = middlePointer + 1
		}

		if nums[middlePointer] < nums[rightPointer] {
			rightPointer = middlePointer - 1
		}

	}

	return 0
}
