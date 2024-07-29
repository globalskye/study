package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{1, 1, 0, 0}))
}
func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}
	end := len(nums) - 1
	currentPropast := 0
	for i := len(nums) - 2; i >= 0; i-- {

		if nums[i] == 0 {
			currentPropast++
			continue
		}
		if nums[i] <= currentPropast {
			if nums[i]-currentPropast >= end {
				return true
			}
			currentPropast++
			continue
		}
		currentPropast = 0
	}
	if currentPropast == 0 {
		return true
	}
	return false
}
