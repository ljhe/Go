package main

import "fmt"

type MinStack struct {
	stack []int
	min   []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		stack: nil,
		min: nil,
	}
	return stack
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || this.min[len(this.min) - 1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min) - 1])
	}
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[0:len(this.stack) - 1]
	this.min = this.min[0:len(this.min) - 1]
}


func (this *MinStack) Top() int {
	s := this.stack[len(this.stack) - 1]
	return s
}


func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min) - 1]
}

func main()  {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param := obj.GetMin()
	fmt.Println(param)
}