package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
func insertToTree(t *TreeNode, val int) {
	if t == nil {
		return
	}
	if t.Val == val {

		return
	}

	if t.Val > val {
		if t.Left == nil {
			t.Left = &TreeNode{Val: val}
			return
		}
		insertToTree(t.Left, val)
		return
	}

	if t.Val < val {
		if t.Right == nil {
			t.Right = &TreeNode{Val: val}
			return
		}
		insertToTree(t.Right, val)
		return
	}

	return

}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return &TreeNode{}
	}
	root := &TreeNode{Val: nums[0]}
	left := &TreeNode{}
	right := &TreeNode{}
	for i := 0; i < len(nums)/2; i++ {
		fmt.Println(nums[len(nums)/2-1-i])
		fmt.Println(nums[len(nums)/2+1+i])
		insertToTree(left, nums[len(nums)/2-1-i])
		insertToTree(right, nums[len(nums)/2+1+i])
	}

	root.Left = left
	root.Right = right
	return root
}
