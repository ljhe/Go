package main

type MyCircularQueue struct {
	size int
	queue []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	myCircularQueue := MyCircularQueue{
		size:  k,
		queue: nil,
	}
	return myCircularQueue
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if !this.IsFull() {
		this.queue = append(this.queue, value)
		return true
	}
	return false
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
		this.queue = this.queue[1:]
		return true
	}
	return false
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if !this.IsEmpty() {
		q := this.queue[0]
		return q
	}
	return -1
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		q := this.queue[len(this.queue) - 1]
		return q
	}
	return -1
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if len(this.queue) >= this.size {
		return true
	}
	return false
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
