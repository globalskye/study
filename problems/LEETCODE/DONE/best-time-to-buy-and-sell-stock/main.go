package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{2, 1, 4}))
}
func maxProfit(prices []int) int {
	i := 0
	j := i + 1
	m := 0
	for j < len(prices) {
		currentProfit := prices[j] - prices[i]
		if currentProfit > m {
			m = currentProfit
		}
		if prices[i] > prices[j] {
			i = j
		}
		j++
	}
	return m
}
