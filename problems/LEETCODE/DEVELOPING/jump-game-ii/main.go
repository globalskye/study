package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
func jump(nums []int) int {
	i := 0
	j := 1
	m := 0

	res := 0
	for {
		if nums[j] > m {
			m = nums[j]
			res++
		}
		j++
	}

}
