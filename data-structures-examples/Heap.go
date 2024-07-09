package main


import (
	"container/heap"
	"fmt"
)

type IntHeap struct {
	arr []int
}

func (h IntHeap) Len() int           { return len(h.arr) }
func (h IntHeap) Less(i, j int) bool { return h.arr[i] > h.arr[j] }
func (h IntHeap) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func (h *IntHeap) Push(x any) {
	h.arr = append(h.arr, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[0 : n-1]
	return x
}
func main() {
	h := &IntHeap{[]int{1, 2, 3, 4, 5}}
	heap.Init(h)
	heap.Push(h, 3)

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
>>>>>>> f79c7b7049bd84938a1bbec433a5526d309c21af
}
