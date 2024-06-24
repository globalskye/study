package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(addTwoNumbers(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, nil))
	l := &ListNode{}

	fmt.Println(l.Val)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var q *ListNode
	var dec int

	for {

		var val2 int
		var val1 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		r := (val1 + val2 + dec) % 10 // 9+9
		dec = (val1 + val2 + dec) / 10
		cur := &ListNode{Val: r}
		if res == nil {
			res = cur
			q = cur
		} else {
			q.Next = cur
			q = cur
			//cur.Next = res
			//res = cur
		}

		if l1 == nil && l2 == nil {
			if dec == 1 {
				q.Next = &ListNode{Val: 1}
			}
			break
		}

	}
	return res
}
