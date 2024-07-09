package main

import (
	"errors"
	"fmt"
	"strings"
)

// TreeNode data structure represents a typical binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTreeNode(t *TreeNode) {
	if t != nil {
		return
	}
	PrintTreeNode(t)
	fmt.Print("[", t.Val, " ", "]")
	PrintTreeNode(t)
}
func PrintTree(node *TreeNode, prefix string, isTail bool) {
	if node == nil {
		return
	}

	var sb strings.Builder
	sb.WriteString(prefix)

	if isTail {
		sb.WriteString("└── ")
		prefix += "    "
	} else {
		sb.WriteString("├── ")
		prefix += "│   "
	}

	sb.WriteString(fmt.Sprintf("%d\n", node.Val))
	fmt.Print(sb.String())

	if node.Left != nil || node.Right != nil {
		if node.Left != nil {
			PrintTree(node.Left, prefix, node.Right == nil)
		}
		if node.Right != nil {
			PrintTree(node.Right, prefix, true)
		}
	}
}

//	func createTree(Values []int) *TreeNode {
//		if len(Values) == 0 {
//			return nil
//		}
//		nodes := make([]*TreeNode, len(Values))
//		for i, Val := range Values {
//			if Val != -1 {
//				nodes[i] = &TreeNode{Val: Val}
//			}
//		}
//		for i := 0; i*2+1 < len(Values); i++ {
//			if nodes[i] != nil {
//				if i*2+1 < len(Values) {
//					nodes[i].Left = nodes[i*2+1]
//				}
//				if i*2+2 < len(Values) {
//					nodes[i].Right = nodes[i*2+2]
//				}
//			}
//		}
//		return nodes[0]
//	}
func main() {

	t := &TreeNode{Val: 8}

	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	t.Find(11)

	t.Delete(5)
	t.Delete(7)

	t.PrintInorder()
	fmt.Println("")

	fmt.Println("min is %d", t.FindMin())
	fmt.Println("max is %d", t.FindMax())
}

// PrintInorder prints the elements in order
func (t *TreeNode) PrintInorder() {

	if t == nil {

		return
	}

	t.Left.PrintInorder()
	fmt.Print(t.Val)
	t.Right.PrintInorder()
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) Insert(Value int) error {

	if t == nil {

		return errors.New("Tree is nil")
	}

	if t.Val == Value {

		return errors.New("This node Value already exists")
	}

	if t.Val > Value {

		if t.Left == nil {

			t.Left = &TreeNode{Val: Value}
			return nil
		}

		return t.Left.Insert(Value)
	}

	if t.Val < Value {

		if t.Right == nil {

			t.Right = &TreeNode{Val: Value}
			return nil
		}

		return t.Right.Insert(Value)
	}

	return nil
}

// Find finds the treenode for the given node Val
func (t *TreeNode) Find(Value int) (TreeNode, bool) {

	if t == nil {
		return TreeNode{}, false
	}

	switch {
	case Value == t.Val:
		return *t, true
	case Value < t.Val:
		return t.Left.Find(Value)
	default:
		return t.Right.Find(Value)
	}
}

// Delete removes the Item with Value from the tree
func (t *TreeNode) Delete(Value int) {
	t.remove(Value)
}

func (t *TreeNode) remove(Value int) *TreeNode {

	if t == nil {
		return nil
	}

	if Value < t.Val {
		t.Left = t.Left.remove(Value)
		return t
	}
	if Value > t.Val {
		t.Right = t.Right.remove(Value)
		return t
	}

	if t.Left == nil && t.Right == nil {
		t = nil
		return nil
	}

	if t.Left == nil {
		t = t.Right
		return t
	}
	if t.Right == nil {
		t = t.Left
		return t
	}

	smallestValOnRight := t.Right
	for {
		//find smallest Value on the Right side
		if smallestValOnRight != nil && smallestValOnRight.Left != nil {
			smallestValOnRight = smallestValOnRight.Left
		} else {
			break
		}
	}

	t.Val = smallestValOnRight.Val
	t.Right = t.Right.remove(t.Val)
	return t
}

// FindMax finds the max element in the given BST
func (t *TreeNode) FindMax() int {
	if t.Right == nil {
		return t.Val
	}
	return t.Right.FindMax()
}

// FindMin finds the min element in the given BST
func (t *TreeNode) FindMin() int {
	if t.Left == nil {
		return t.Val
	}
	return t.Left.FindMin()
}
