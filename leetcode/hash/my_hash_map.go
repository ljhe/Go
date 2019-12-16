package main

import "fmt"

type MyHashMap struct {
	hashMap map[int]int
}


/** Initialize your data structure here. */
func Constructor1() MyHashMap {
	var myHashMap MyHashMap
	// 如果不初始化 map 那么就会创建一个 nil map. nil map 不能用来存放键值对
	myHashMap.hashMap = make(map[int]int)
	return myHashMap
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.hashMap[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if _, ok := this.hashMap[key]; ok == true {
		return this.hashMap[key]
	}
	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	delete(this.hashMap, key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main()  {
	obj := Constructor1()
	obj.Put(2,1)
	fmt.Println(obj)
	param2 := obj.Get(2)
	obj.Remove(2)
	fmt.Println(param2)
	fmt.Println(obj)
}
