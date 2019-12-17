package main

import (
	"Grammar/leetcode/midForNode/tree"
	"fmt"
	"math"
)
var result []int

// 中序遍历
func inorderTraversal(root *tree.Node) []int {
	if root != nil {
		inorderTraversal(root.Left)
		result = append(result, root.Val)
		inorderTraversal(root.Right)
	}
	return result
}

// 前序遍历
func preorderTraversal(root *tree.Node) []int {
	if root != nil {
		result = append(result, root.Val)
		preorderTraversal(root.Left)
		preorderTraversal(root.Right)
	}
	return result
}

// 后续遍历
func postorderTraversal(root *tree.Node) []int {
	if root != nil {
		postorderTraversal(root.Left)
		postorderTraversal(root.Right)
		result = append(result, root.Val)
	}
	return result
}

// 给定一个二叉树,找出其最大深度.
func maxDepth(root *tree.Node) int {
	depth := 0
	if root != nil {
		depth = int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1)
	}
	return depth
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
	a := preorderTraversal(&node)
	fmt.Println(a)
}
