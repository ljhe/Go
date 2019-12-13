package main

import "fmt"

type LinkedList struct {
	val int
	next *LinkedList
}

type MyLinkedList struct {
	header *LinkedList
	size int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var myLinkedList = MyLinkedList{
		header: &LinkedList{},
		size:   0,
	}
	return myLinkedList
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	// the index is invalid
	if this.size < index + 1 {
		return -1
	}

	prev := this.header
	for i := 0; i <= index; i++ {
		if i != 0 {
			prev = prev.next
		}
	}
	return prev.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	var cur = LinkedList{
		val:  val,
		next: this.header,
	}
	this.header = &cur
	this.size++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	var cur = LinkedList{
		val:  val,
		next: nil,
	}

	prev := this.header
	for i := 0; i < this.size; i++ {
		if i != 0 {
			prev = prev.next
		}
	}
	prev.next = &cur
	this.size++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	// If index is greater than the length, the node will not be inserted.
	if index > this.size {
		return
	}
	// 头节点添加
	if index == 0 {
		this.AddAtHead(val)
	}
	// 中间节点添加
	if index != 0 && index != this.size {
		prev := this.header
		for i := 0; i < index; i++  {
			if i != 0 {
				prev = prev.next
			}
		}
		prev.next = &LinkedList{
			val:  val,
			next: prev.next,
		}
		this.size++
	}
	// 尾节点添加
	if index == this.size {
		this.AddAtTail(val)
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	// the index is invalid
	if this.size < index + 1 {
		return
	}

	// 删除头节点
	if index == 0 {
		this.header = this.header.next
		this.size--
		return
	}

	prev := this.header
	forDelLast := &LinkedList{}
	for i := 0; i < index; i++ {
		if i != 0 {
			prev = prev.next
		}
		if i + 1 == index {
			forDelLast = prev
		}
	}

	// 删除中间节点
	if index != 0 && index + 1 != this.size {
		prev.next = prev.next.next
		this.size--
		return
	}

	// 删除尾节点
	if index + 1 == this.size {
		forDelLast.next = nil
		this.size--
		return
	}
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main()  {
	obj := Constructor()

	// 顺序 1 > 4 > 2 > 3
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 4)

	obj.DeleteAtIndex(4)
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Printf("%+v",obj)
}
