package main

import (
	"strconv"
)

type Heap struct {
	arr []int
}

func (h *Heap) Heapify(i int) {
	for {
		l := i*2 + 1
		r := i*2 + 2
		largest := i

		if l < len(h.arr) && h.arr[l] > h.arr[largest] {
			largest = l
		}
		if r < len(h.arr) && h.arr[r] > h.arr[largest] {
			largest = r
		}
		if largest == i {
			break
		}

		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		i = largest
	}
}

func BuildHeap(arr []int) *Heap {
	r := &Heap{arr: arr}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		r.Heapify(i)
	}
	return r
}

func (h *Heap) Pop() int {
	n := len(h.arr)
	if n == 0 {
		return 0
	}
	root := h.arr[0]
	h.arr[0] = h.arr[n-1]
	h.arr = h.arr[:n-1]
	h.Heapify(0)
	return root
}

func findRelativeRanksUsingHeap(score []int) []string {
	m := make(map[int]int, len(score))
	for i := 0; i < len(score); i++ {
		m[score[i]] = i
	}

	heap := BuildHeap(score)
	res := make([]string, len(score))

	for i := 0; i < len(score); i++ {
		val := heap.Pop()
		idx := m[val]
		switch i {
		case 0:
			res[idx] = "Gold Medal"
		case 1:
			res[idx] = "Silver Medal"
		case 2:
			res[idx] = "Bronze Medal"
		default:
			res[idx] = strconv.Itoa(i + 1)
		}
	}

	return res
}
