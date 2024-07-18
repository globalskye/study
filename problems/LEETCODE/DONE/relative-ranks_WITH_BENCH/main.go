package main

import "fmt"

func main() {
	//go test -bench=.
	arr := []int{10, 3, 8, 9, 4}
	fmt.Println("Using Sort:", findRelativeRanksUsingSort(arr))
	fmt.Println("Using Heap:", findRelativeRanksUsingHeap(arr))
}
