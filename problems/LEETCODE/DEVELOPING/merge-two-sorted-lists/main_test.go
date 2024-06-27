package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			list1:    &ListNode{Val: 1, Next: nil},
			list2:    nil,
			expected: &ListNode{Val: 1, Next: nil},
		},
		{
			list1:    nil,
			list2:    &ListNode{Val: 1, Next: nil},
			expected: &ListNode{Val: 1, Next: nil},
		},
		{
			list1:    &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}},
			list2:    &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
		},
		{
			list1:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
			list2:    &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
		},
		{
			list1:    &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}},
			list2:    &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: nil}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}},
		},
	}

	for _, test := range tests {
		result := mergeTwoLists(test.list1, test.list2)
		if !isEqual(result, test.expected) {
			t.Errorf("mergeTwoLists() = %v, want %v", listToSlice(result), listToSlice(test.expected))
		}
	}
}

func isEqual(list1 *ListNode, list2 *ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	return list1 == nil && list2 == nil
}

func listToSlice(list *ListNode) []int {
	var result []int
	for list != nil {
		result = append(result, list.Val)
		list = list.Next
	}
	return result
}

func main() {
	// The main function is just a placeholder in this context.
	// The real entry point for running tests is `go test`.
	fmt.Println("Run `go test` to execute the tests")
}
