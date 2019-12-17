package my_hash

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

// 给定一个非空整数数组,除了某个元素只出现一次以外,其余每个元素均出现两次,找出那个只出现了一次的元素.
func (this *MyHashMap) singleNumber(nums []int) int {
	for _, v := range nums {
		if this.Get(v) != -1 {
			delete(this.hashMap, v)
		}
	}
	return nums[0]
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
	fmt.Println(obj)
}
