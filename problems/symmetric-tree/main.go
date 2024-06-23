package main

import "fmt"

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}

	fmt.Println(isSymmetric(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func GetValueInorder(t *TreeNode, stack *[]int) {
//
//		if t == nil {
//
//			return
//		}
//
//		GetValueInorder(t.Left, stack)
//		*stack = append(*stack, t.Val)
//		GetValueInorder(t.Right, stack)
//	}
func test(t1 *TreeNode, t2 *TreeNode) {

	fmt.Printf("%v : %v", t1.Val, t2.Val)
	test(t1.Left, t2.Right)
}

//func GetValueInorder1(t *TreeNode, stack *[]int) {
//	if t == nil {
//		return
//	}
//	GetValueInorder1(t.Left, stack)
//	*stack = append(*stack, t.Val)
//	if t.Left != nil && t.Right == nil {
//		*stack = append(*stack, 1000)
//	}
//	GetValueInorder1(t.Right, stack)
//}
//
//func GetValueInorder2(t *TreeNode, stack *[]int) {
//	if t == nil {
//		return
//	}
//	GetValueInorder2(t.Right, stack)
//	*stack = append(*stack, t.Val)
//	if t.Right != nil && t.Left == nil {
//		*stack = append(*stack, 1000)
//	}
//	GetValueInorder2(t.Left, stack)
//}

func isSymmetric(root *TreeNode) bool {
	test(root.Left, root.Right)
	return true
	//if root == nil {
	//	return true
	//}
	//
	//res1 := make([]int, 0)
	//res2 := make([]int, 0)
	//
	//GetValueInorder1(root.Left, &res1)
	//GetValueInorder2(root.Right, &res2)
	//fmt.Println(res1)
	//fmt.Println(res2)
	//if len(res1) != len(res2) {
	//	return false
	//}
	//for i := 0; i < len(res1); i++ {
	//	if res1[i] != res2[i] {
	//		return false
	//	}
	//}
	//
	//return true
}
