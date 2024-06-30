package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}}
	deleteDuplicates(node)
}
func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{Val: head.Val}
	tmp := res
	for head != nil {

		if head.Val == tmp.Val {
			head = head.Next
			continue
		}

		tmp.Next = &ListNode{Val: head.Val}
		tmp = tmp.Next

		head = head.Next
	}

	return res
}
