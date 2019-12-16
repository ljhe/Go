package main

import "Grammar/leetcode/hash/my_hash"

func containsDuplicate(nums []int) bool {
	i := 0
	myHash := my_hash.Constructor()
	for _, v := range nums {
		if myHash.AddForContainsDuplicate(v) == false {
			i++
		}
	}
	if i >= 1 {
		return true
	}
	return false
}

func main()  {
	
}
