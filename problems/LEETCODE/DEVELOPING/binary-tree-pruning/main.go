package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val == 0 {
			if node.Right != nil && node.Right.Val == 0 {
				node.Right = nil
			}
			if node.Left != nil && node.Left.Val == 0 {
				node.Left = nil
			}
		}
		if node.Right != nil {

			queue = append(queue, node.Right)
		}
		if node.Left != nil {

			queue = append(queue, node.Left)

		}
	}
	return root
}
