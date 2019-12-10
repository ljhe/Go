package main

import "fmt"

type MyStack struct {
	queue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	var stack MyStack
	return stack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if !this.Empty() {
		q := this.queue[len(this.queue) - 1]
		this.queue = this.queue[0 : len(this.queue) - 1]
		return q
	}
	return 0
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if !this.Empty() {
		q := this.queue[len(this.queue) - 1]
		return q
	}
	return 0
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func main()  {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	param2 := obj.Pop()
	param3 := obj.Top()
	param4 := obj.Empty()
	fmt.Println(param2)
	fmt.Println(param3)
	fmt.Println(param4)
}
