package main

import (
	"Grammar/leetcode/midForNode/tree"
	"fmt"
)
var result []int

func inorderTraversal(root *tree.Node) []int {
	if root != nil {
		inorderTraversal(root.Left)
		result = append(result, root.Val)
		inorderTraversal(root.Right)
	}
	return result
}

func main()  {
	var node tree.Node
	node = tree.Node{}
	node.Left = &tree.Node{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node.Left.Left = new(tree.Node)
	node.Left.Left.SetValue(3)
	node.Left.Right = new(tree.Node)
	node.Left.Right.SetValue(4)
	a := inorderTraversal(&node)
	fmt.Println(a)
}
