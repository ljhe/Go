package main

/**
 * 项目路径 F:\Code\Git\Go\src\Grammar
 * GOPATH路径 F:\Code\Git\Go
 * import路径 Grammar/tree
 */
import "Grammar/tree"

/**
 * 对现有结构体的扩展
 */
type myTreeNode struct {
	node *tree.Node
}

/**
 * 后续遍历
 * 输出 2 0 4 5 3
 */
func (myNode *myTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

/**
 * 创建了一个这样的树
 * 				  3
 *			0			5
 *				2	4
 */
func main()  {
	var root tree.Node
	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{
		Value: 5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	root.Traverse()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}