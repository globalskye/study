package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}

	result := mergeTwoLists(list1, list2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}

	tail := result

	if list1 == nil && list2 == nil {
		return result
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}
	for list1 != nil && list2 != nil {

		if list1.Val >= list2.Val {

			tail.Next = &ListNode{Val: list2.Val}
			tail = tail.Next

			list2 = list2.Next
			continue
		}
		if list2.Val >= list1.Val {

			tail.Next = &ListNode{Val: list1.Val}
			tail = tail.Next

			list1 = list1.Next

		}
	}
	if list1 == nil {
		tail.Next = list2

	}
	if list2 == nil {
		tail.Next = list1

	}
	return result.Next
}
