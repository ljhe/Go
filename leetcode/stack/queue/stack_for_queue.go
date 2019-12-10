package main

import "fmt"

type MyQueue struct {
	stack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	var q MyQueue
	q = MyQueue{stack:nil}
	return q
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stack = append(this.stack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if ! this.Empty() {
		q := this.stack[0]
		this.stack = this.stack[1:]
		return q
	}
	return 0
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if ! this.Empty() {
		q := this.stack[0]
		return q
	}
	return 0
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stack) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main()  {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj)
	param2 := obj.Pop()
	param3 := obj.Peek()
	param4 := obj.Empty()
	fmt.Println(param2)
	fmt.Println(param3)
	fmt.Println(param4)
}