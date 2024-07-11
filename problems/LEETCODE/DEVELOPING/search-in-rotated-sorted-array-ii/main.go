package main

import (
	"fmt"
)

func main() {

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
func binarySearch(nums []int, target int) int {
	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer <= rightPointer {
		middlePointer := (leftPointer + rightPointer) / 2
		currentValue := nums[middlePointer]

		if currentValue == target {
			return middlePointer
		}

		if nums[leftPointer] < target {

			leftPointer = middlePointer + 1
		}

		if nums[rightPointer] > target {
			rightPointer = middlePointer - 1
		}

	}
	return -1
}
func search(nums []int, target int) int {

	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer <= rightPointer {
		middlePointer := (leftPointer + rightPointer) / 2
		if target == nums[middlePointer] {
			return middlePointer
		}
		if nums[middlePointer] > nums[middlePointer+1] {
			if nums[0] < target {

				return search(nums[middlePointer+1:], target)
			}
			leftPointer = middlePointer + 1
			rightPointer = len(nums) - 1
			return search(nums[middlePointer+1:], target)
		}
		if nums[middlePointer] < nums[middlePointer-1] {

			middlePointerT := (leftPointer+rightPointer)/2 - 1
			if target == nums[middlePointerT] {
				return middlePointerT
			}
			if nums[middlePointerT] > nums[leftPointer] {
				leftPointer = middlePointerT + 1
			}

			if nums[middlePointerT] < nums[rightPointer] {
				rightPointer = middlePointerT - 1
			}

			continue
		}
		if nums[middlePointer] < nums[leftPointer] {
			leftPointer = middlePointer + 1
		}

		if nums[middlePointer] > nums[rightPointer] {
			rightPointer = middlePointer - 1
		}

	}

	return -1
}
