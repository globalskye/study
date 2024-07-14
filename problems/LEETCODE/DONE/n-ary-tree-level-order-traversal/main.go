package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	queue := []*Node{root}
	res := [][]int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		currentValue := []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentValue = append(currentValue, node.Val)
			if len(node.Children) > 0 {
				for j := 0; j < len(node.Children); j++ {
					queue = append(queue, node.Children[j])
				}
			}
			res = append(res, currentValue)
		}
	}
	return res
}
