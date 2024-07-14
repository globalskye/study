package main

import "fmt"

func main() {
	tree := createTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(levelOrderBottom(tree))
}
func createTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for index < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if vals[index] != nil {
			leftVal := vals[index].(int)
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}
		index++
		if index < len(vals) && vals[index] != nil {
			rightVal := vals[index].(int)
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	currentArr := []int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		currentArr = []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentArr = append(currentArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append([][]int{currentArr}, res...)
	}
	return res
}
