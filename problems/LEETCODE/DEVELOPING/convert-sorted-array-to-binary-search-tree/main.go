package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(node *TreeNode, space int, level int) {
	if node == nil {
		return
	}

	// Увеличиваем расстояние между уровнями
	space += level

	// Печатаем правого ребенка сначала
	printTree(node.Right, space, level)

	// Печатаем текущее узел после пробелов
	fmt.Print("\n")
	for i := level; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Print(node.Val, "\n")

	// Печатаем левого ребенка
	printTree(node.Left, space, level)
}
func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	printTree(root, 0, 10)
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
	root := &TreeNode{}
	j := 0
	i := len(nums) - 1
	if len(nums)%2 == 1 {
		root = &TreeNode{Val: nums[len(nums)/2]}
	}

	for i < len(nums) {

		insertToTree(root.Left, nums[i])
		insertToTree(root.Right, nums[j])
		i--
		j++

	}

	return root
}
