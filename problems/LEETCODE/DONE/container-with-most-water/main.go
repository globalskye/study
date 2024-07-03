package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 0, 0, 0, 0, 0, 0, 2, 2}))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxArea(height []int) int {
	area := 0

	i := 0
	j := len(height) - 1
	for i < j {
		currentHeight := Min(height[j], height[i])

		currentWeight := j - i

		if currentWeight*currentHeight > area {
			area = currentHeight * currentWeight
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return area
}
