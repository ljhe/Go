package tree

import "fmt"

type Node struct {
	Val int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Val: value}
}

func (node *Node) SetValue (value int)  {
	// nil 指针也可以调用方法 所以这里加了判断
	if node == nil {
		fmt.Println("Setting Value to nil Node. Ignored")
		return
	}
	node.Val = value
}