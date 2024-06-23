package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}

	// Создаем второй список list2 = [1, 3, 4]
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	result := mergeTwoLists(list1, list2)
	for result.Next != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}

type LinkedList struct {
	head *ListNode
}

func (list *LinkedList) addToEnd(val int) {
	newNode := &ListNode{Val: val, Next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode

}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result LinkedList

	if list1 == nil && list2 == nil {
		return result.head
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}
	for {
		val1 := list1.Val
		val2 := list2.Val

		if val1 == val2 {

			result.addToEnd(val2)
			result.addToEnd(val1)

		}
		if val1 > val2 {
			result.addToEnd(val2)
			result.addToEnd(val1)
		}
		if val2 > val1 {
			result.addToEnd(val1)
			result.addToEnd(val2)
		}

		if list1.Next == nil && list2.Next == nil {
			break
		}
		if list1.Next != nil {
			list1 = list1.Next
		}
		if list2.Next != nil {
			list2 = list2.Next
		}

	}

	return result.head
}
