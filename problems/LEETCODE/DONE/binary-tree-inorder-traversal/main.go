package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetValueInorder(t *TreeNode, stack *[]int) {

	if t == nil {
		return
	}

	GetValueInorder(t.Left, stack)
	*stack = append(*stack, t.Val)
	GetValueInorder(t.Right, stack)

}
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	GetValueInorder(root, &res)

	return res
}
