package main

import (
	"fmt"
)

func createTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(values))
	for i, val := range values {
		if val != -1 {
			nodes[i] = &TreeNode{Val: val}
		}
	}
	for i := 0; i*2+1 < len(values); i++ {
		if nodes[i] != nil {
			if i*2+1 < len(values) {
				nodes[i].Left = nodes[i*2+1]
			}
			if i*2+2 < len(values) {
				nodes[i].Right = nodes[i*2+2]
			}
		}
	}
	return nodes[0]
}

func main() {

	values1 := []int{1, 2, 2, 3, 4, 4, 3}
	values2 := []int{1, 2, 2, 3, 4, 4, 3}
	values3 := []int{1, 2, 2, 3, 4, 4, 3}
	root1 := createTree(values1)
	root2 := createTree(values2)
	root3 := createTree(values3)
	fmt.Println(isSymmetric(root1))
	fmt.Println(isSymmetric(root2))
	fmt.Println(isSymmetric(root3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetValueInorder(t1 *TreeNode, t2 *TreeNode, stack *[]bool) {
	if t1 == nil && t2 != nil {
		*stack = append(*stack, false)
		return
	}
	if t2 == nil && t1 != nil {
		*stack = append(*stack, false)
		return
	}
	if t1 == nil && t2 == nil {
		*stack = append(*stack, true)
		return
	}
	if t1.Val != t2.Val {
		*stack = append(*stack, false)
		return
	}
	GetValueInorder(t1.Left, t2.Right, stack)

	GetValueInorder(t1.Right, t2.Left, stack)

}
func isSymmetric(root *TreeNode) bool {
	stack := make([]bool, 0)
	GetValueInorder(root.Left, root.Right, &stack)
	for _, ok := range stack {
		if !ok {
			return false
		}
	}

	return true

}
