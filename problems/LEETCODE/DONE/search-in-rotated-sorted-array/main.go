package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchDiv([]int{3, 5, 1}))
	fmt.Println(search([]int{3, 5, 1}, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if r := searchDiv(nums); r == -1 {
		return binarySearch(nums, target)
	}
	div := searchDiv(nums)
	if target < nums[0] {
		if r := binarySearch(nums[div:], target); r != -1 {
			return div + r
		}
		return -1
	}
	if r := binarySearch(nums[:div], target); r != -1 {
		return r
	}
	return -1

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

		if nums[middlePointer] < target {

			leftPointer = middlePointer + 1
		}

		if nums[middlePointer] > target {
			rightPointer = middlePointer - 1
		}

	}
	return -1
}
func searchDiv(nums []int) int {

	leftPointer := 0
	rightPointer := len(nums) - 1
	if nums[leftPointer] < nums[rightPointer] {
		return -1
	}
	for leftPointer <= rightPointer {
		middlePointer := (leftPointer + rightPointer) / 2

		if middlePointer < rightPointer && nums[middlePointer] > nums[middlePointer+1] {
			return middlePointer
		}
		if middlePointer > leftPointer && nums[middlePointer] < nums[middlePointer-1] {
			return middlePointer
		}

		if nums[middlePointer] < nums[leftPointer] {
			rightPointer = middlePointer - 1
			continue
		}
		if nums[middlePointer] > nums[rightPointer] {
			leftPointer = middlePointer + 1
		}

	}

	return -1
}
