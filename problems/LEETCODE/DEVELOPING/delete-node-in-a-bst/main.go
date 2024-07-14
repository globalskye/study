package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func main() {
	//a := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{}}, Right: &TreeNode{Val: 6}}
	b := createTree([]interface{}{0})
	b = deleteNode(b, 0)
	fmt.Println(b)
}

func deleteNode(root *TreeNode, key int) *TreeNode {

	return remove(root, key)
}
func remove(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > target {
		node := remove(root.Left, target)
		root.Left = node
		return root
	}
	if root.Val < target {
		node := remove(root.Right, target)
		root.Right = node
		return root

	}

	if root.Right != nil {
		n := &TreeNode{}
		n = root.Right
		n.Left = root.Left

		return n
	}
	if root.Left != nil {
		n := &TreeNode{}
		n = root.Left
		return n
	}
	root = nil
	return root
}
