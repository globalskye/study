package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(Values []int) *TreeNode {
	if len(Values) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(Values))
	for i, Val := range Values {
		if Val != -1 {
			nodes[i] = &TreeNode{Val: Val}
		}
	}
	for i := 0; i*2+1 < len(Values); i++ {
		if nodes[i] != nil {
			if i*2+1 < len(Values) {
				nodes[i].Left = nodes[i*2+1]
			}
			if i*2+2 < len(Values) {
				nodes[i].Right = nodes[i*2+2]
			}
		}
	}
	return nodes[0]
}
func isValidBST(root *TreeNode) bool {
	arr := []int{}
	getArrayFromBST(&arr, root)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
func getArrayFromBST(arr *[]int, root *TreeNode) {

	if root == nil {
		return
	}
	getArrayFromBST(arr, root.Left)
	*arr = append(*arr, root.Val)
	getArrayFromBST(arr, root.Right)

}
func main() {
	root := createTree([]int{2, 1, 3})

	arr := []int{}
	getArrayFromBST(&arr, root)
	fmt.Println(arr)
}
