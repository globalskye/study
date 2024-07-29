package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 100, 10}))
}
func maxProfit(prices []int) int {
	profits := make([]int, 0)
	i := 0
	j := 1
	res := 0
	for j < len(prices) {
		currentProfit := prices[j] - prices[i]
		if currentProfit > 0 {
			profits = append(profits, currentProfit)
			i++
		}
		if prices[i] >= prices[j] {
			i = j
		}
		j++
	}
	for i = 0; i < len(profits); i++ {
		res = res + profits[i]
	}
	return res
}
