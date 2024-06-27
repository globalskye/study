package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}

	// Создаем второй список list2 = [1, 3, 4]
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}

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
	result := &ListNode{}

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
		if list1.Val == list2.Val {
			fmt.Println(list1.Val)
			fmt.Println(list2.Val)
			list1 = list1.Next
			list2 = list2.Next
		}
		if list2 != nil {
			curList := list2
			for list2.Val >= list1.Val {
				if list2.Next == nil {
					break
				}
				list2 = list2.Next

			}
			fmt.Println(list1.Val)
			fmt.Println(list2.Val)

			list2 = curList
			list2 = list2.Next
			continue
		}
		if list1 != nil {
			curList := list1
			for list1.Val >= list2.Val {
				list2 = list2.Next
			}
			fmt.Println(list2.Val)
			fmt.Println(list1.Val)

			list1 = curList
			list1 = list1.Next
			continue
		}

	}

	return result
}
