package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 :=
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	reverseBetween(list1, 2, 4)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	i := 1
	j := 1
	buff1 := 0
	buff2 := 0
	var tmp ListNode

	fmt.Println(buff2)
	fmt.Println(buff1)
	for head != nil {
		if i == left {

			buff1 = head.Val
			tmp = *head
		}
		if j == right {

			buff2 = head.Val

			head.Val = buff1
			break
		}
		i++
		j++

	}
	return &tmp
}
