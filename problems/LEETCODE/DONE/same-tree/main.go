package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	tree1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Дерево 2
	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}
	fmt.Println(isSameTree(tree1, tree2))
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q == nil {
		return false
	}
	if q != nil && p == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}
