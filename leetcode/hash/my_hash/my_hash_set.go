package my_hash

import "fmt"

type MyHashSet struct {
	value map[int][]int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	var myHashSet MyHashSet
	myHashSet.value = make(map[int][]int)
	return myHashSet
}


func (this *MyHashSet) Add(key int)  {
	k := this.HashFunc(key)
	flag := true
	if len(this.value[k]) == 0 {
		this.value[k] = append(this.value[k], key)
	} else {
		for _, val := range this.value[k] {
			if val == key {
				flag = false
				return
			}
		}
		if flag == true {
			this.value[k] = append(this.value[k], key)
		}
	}
}

func (this *MyHashSet) AddForContainsDuplicate(key int) bool {
	k := this.HashFunc(key)
	flag := true
	if len(this.value[k]) == 0 {
		this.value[k] = append(this.value[k], key)
	} else {
		for _, val := range this.value[k] {
			if val == key {
				flag = false
				break
			}
		}
		if flag == true {
			this.value[k] = append(this.value[k], key)
		}
	}
	return flag
}

func (this *MyHashSet) Remove(key int)  {
	k := this.HashFunc(key)
	if len(this.value[k]) != 0 {
		for index, val := range this.value[k] {
			if val == key {
				this.value[k] = append(this.value[k][:index], this.value[k][index + 1:]...)
			}
		}
	}
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	flag := false
	k := this.HashFunc(key)
	for _, v := range this.value[k] {
		if v == key {
			flag = true
		}
	}
	return flag
}

func (this *MyHashSet) HashFunc(key int) int {
	return key % 5
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main()  {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	obj.Add(2)
	fmt.Printf("%+v", obj)
	obj.Remove(2)
}
