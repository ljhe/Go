package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	// 最大值
	var max = 0
	// 最大值索引
	var key = 0
	for k, v := range nums {
		if max < v {
			max = v
			key = k
		}
	}
	for k, v := range nums{
		if max < v * 2 && k != key {
			return -1
		}
	}
	return key
}

func main()  {
	var nums []int
	nums = []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}
