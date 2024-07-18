package main

import "fmt"

func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	current := root
	for i := 1; i < len(values); i++ {
		if values[i] != nil {
			current.Right = &TreeNode{Val: values[i].(int)}
			current = current.Right
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	values := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := createTree(values)
	fmt.Println(minDepth(root))
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	res := 1
	for len(queue) > 0 {

		root = queue[0]
		queue = queue[1:]
		if root.Left == nil && root.Right == nil {
			return res
		}
		if root.Right == nil && root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				res++
				return res
			}
			res++

		}
		if root.Left == nil && root.Right != nil {
			if root.Right.Left == nil && root.Right.Right == nil {
				res++
				return res
			}
			res++
		}
		if root.Left != nil && root.Right != nil {
			res++
		}
		if root.Left != nil && root.Right != nil {
			res++
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}

	}
	return res
}
func goingDeeper(a *int, root *TreeNode) {
	if root == nil {
		return
	}
	*a--
	goingDeeper(a, root.Left)
	goingDeeper(a, root.Right)
	*a++
}
