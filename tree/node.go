package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

/**
 * 工厂函数
 * 这里返回了一个局部变量的地址给外面用, go 语言特性
 */
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

/**
 * set 方法这里要的是一个指针接收者 只有使用指针才可以改变结构内容, 但是下面调用的 root 还是值类型
 * 值接收者和指针接收者非常灵活, go 语言会解析出来它们需要的
 * 值接收者 会把值拷贝一份传过来
 * 指针接收者 会把地址解析出来传过来
 * nil 指针也可以调用方法
 */
func (node Node) Print ()  {
	fmt.Println(node.Value)
}

func (node *Node) SetValue (value int)  {
	// nil 指针也可以调用方法 所以这里加了判断
	if node == nil {
		fmt.Println("Setting Value to nil Node. Ignored")
		return
	}
	node.Value = value
}


