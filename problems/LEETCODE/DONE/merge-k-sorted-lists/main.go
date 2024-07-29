package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type Heap struct {
	arr []int
}

// [[1,4,5],[1,3,4],[2,6]]
func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}
	r := []*ListNode{list1, list2, list3}
	fmt.Println(mergeKLists(r))
}
func (h *Heap) Pop() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}
	tmp := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.Heapify(0)
	return tmp, true
}
func (h *Heap) Heapify(i int) {

	for {
		l := i*2 + 1
		r := i*2 + 2
		largest := i
		if l < len(h.arr) && h.arr[l] < h.arr[largest] {
			largest = l
		}
		if r < len(h.arr) && h.arr[r] < h.arr[largest] {
			largest = r
		}
		if i == largest {
			break
		}
		h.arr[largest], h.arr[i] = h.arr[i], h.arr[largest]
		i = largest

	}
}

// facepalm :)
func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}

	heap := Heap{arr: []int{}}
	for _, v := range lists {
		for v != nil {

			heap.arr = append([]int{v.Val}, heap.arr...)
			heap.Heapify(0)
			v = v.Next
		}
	}

	tmp := res

	for len(heap.arr) > 0 {
		val, _ := heap.Pop()
		res.Next = &ListNode{Val: val}
		res = res.Next
	}

	return tmp.Next
}
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}