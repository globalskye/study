package main

import (
	"fmt"
)

type BST struct {
	Val       int
	LeftNode  *BST
	RightNode *BST
}

func (t *BST) Insert(val int) {
	if t == nil {

		return
	}

	if val < t.Val {
		if t.LeftNode == nil {
			t.LeftNode = &BST{Val: val}
		} else {
			t.LeftNode.Insert(val)
		}
	} else {
		if t.RightNode == nil {
			t.RightNode = &BST{Val: val}
		} else {
			t.RightNode.Insert(val)
		}
	}
}
func BuildTree(nums []int) BST {
	if len(nums) == 0 {
		return BST{}
	}
	node := BST{}
	for i := 0; i < len(nums); i++ {
		node.Insert(nums[i])
	}
	return node
}
func (t *BST) PrintPreOrder() {
	if t == nil {

		return
	}
	fmt.Println(t.Val)
	t.LeftNode.PrintPreOrder()
	t.RightNode.PrintPreOrder()
}

func (t *BST) PrintPostOrder() {
	if t == nil {

		return
	}

	t.LeftNode.PrintPreOrder()
	t.RightNode.PrintPreOrder()
	fmt.Println(t.Val)
}
func (t *BST) PrintInOrder() {
	if t == nil {

		return
	}
	t.LeftNode.PrintPreOrder()
	fmt.Println(t.Val)
	t.RightNode.PrintPreOrder()
}

func (t *BST) Find(val int) (BST, bool) {
	if t == nil {
		return BST{}, false
	}
	if t.Val == val {
		return *t, true
	}
	if t.Val < val {
		return t.LeftNode.Find(val)
	} else {
		return t.RightNode.Find(val)
	}
}
func (t *BST) Delete(val int) {
	t.remove(val)
}
func (t *BST) remove(val int) *BST {
	if t == nil {
		return nil
	}

	if val < t.Val {
		t.LeftNode = t.LeftNode.remove(val)
		return t
	}
	if val > t.Val {
		t.RightNode = t.RightNode.remove(val)
		return t
	}

	if t.LeftNode == nil && t.RightNode == nil {
		t = nil
		return nil
	}

	if t.LeftNode == nil {
		t = t.RightNode
		return t
	}
	if t.RightNode == nil {
		t = t.LeftNode
		return t
	}

	smallestValOnRight := t.RightNode
	for {
		//find smallest Value on the Right side
		if smallestValOnRight != nil && smallestValOnRight.LeftNode != nil {
			smallestValOnRight = smallestValOnRight.LeftNode
		} else {
			break
		}
	}

	t.Val = smallestValOnRight.Val
	t.RightNode = t.RightNode.remove(t.Val)
	return t
}
func main() {
	t := &BST{}
	t.Val = 1
	t.Insert(5)
	t.Insert(4)
	t.Insert(6)
	t.Insert(2)
	t.Insert(1)
	t.Insert(3)

}
func pop(h *[]int) int {

	return 0
}
