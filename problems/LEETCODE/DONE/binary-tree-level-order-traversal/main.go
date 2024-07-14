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
func (t *TreeNode) PrintPreOrder() {
	if t == nil {

		return
	}
	fmt.Println(t.Val)
	t.Left.PrintPreOrder()
	t.Right.PrintPreOrder()
}

func (t *TreeNode) PrintPostOrder() {
	if t == nil {

		return
	}
	fmt.Println(t.Val)
	t.Left.PrintPreOrder()
	t.Right.PrintPreOrder()

}
func (t *TreeNode) PrintInOrder() {
	if t == nil {
		return
	}
	t.Left.PrintPreOrder()
	fmt.Println(t.Val)
	t.Right.PrintPreOrder()
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

func main() {
	tree := createTree([]interface{}{3, 9, 20, 15, 7, 1, 3, 1})
	levelOrder(tree)

}
