package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	root := createTree([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(maxDepth(root))
}
func GoingDeeper(t *TreeNode, count *int, max *int) {
	//s := t
	//fmt.Println(s)
	buf := *count
	if t == nil {
		if *max >= buf {
			return
		}
		*max = buf

		return
	}
	*count++
	GoingDeeper(t.Left, count, max)

	GoingDeeper(t.Right, count, max)
	*count--
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	max := 0
	GoingDeeper(root, &count, &max)

	return max
}
func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}
