package main

import "fmt"

func main() {
	fmt.Println(searchDiv([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}))
	fmt.Println(searchDiv([]int{1, 0, 1, 1, 1}))
	fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2))
}
func search(nums []int, target int) bool {
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	div := searchDiv(nums)
	if div == -1 {
		if r := binarySearch(nums, target); r == -1 {
			return false
		}
		return true
	}
	if target < nums[0] {
		if r := binarySearch(nums[div:], target); r != -1 {
			return true
		}
		return false
	}
	if r := binarySearch(nums[:div], target); r != -1 {
		return true
	}
	return false
}
func binarySearch(nums []int, target int) int {
	leftPointer := 0
	rightPointer := len(nums) - 1
	for leftPointer <= rightPointer {
		middlePointer := (leftPointer + rightPointer) / 2

		if nums[middlePointer] == target {
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
	d := -1
	if nums[leftPointer] < nums[rightPointer] {
		return -1
	}

	for leftPointer <= rightPointer {
		middlePointer := (leftPointer + rightPointer) / 2

		if nums[leftPointer] == nums[middlePointer] && nums[middlePointer] == nums[rightPointer] {
			leftPointer++
			rightPointer--
			d++
			continue
		}

		if middlePointer > leftPointer && nums[middlePointer] < nums[middlePointer-1] {
			return middlePointer
		}

		if middlePointer < rightPointer && nums[middlePointer] > nums[middlePointer+1] {
			return middlePointer + 1
		}

		if nums[leftPointer] < nums[middlePointer] {
			leftPointer = middlePointer + 1

		} else {
			rightPointer = middlePointer - 1

		}
		d--
	}
	if d != -1 {
		return d
	}

	return -1
}
