package main

import "fmt"

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6}, 6))
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

		if nums[leftPointer] < target {

			leftPointer = middlePointer + 1
		}

		if nums[rightPointer] > target {
			rightPointer = middlePointer - 1
		}

	}

	return 0

}
