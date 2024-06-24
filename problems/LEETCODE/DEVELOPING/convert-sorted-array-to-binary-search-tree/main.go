package main

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
	for i := 0; i < len(nums); i++ {
		insertToTree(root, nums[i])
	}
	return root
}
