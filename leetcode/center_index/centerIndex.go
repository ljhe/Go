package main

import "fmt"

func pivotIndex(nums []int) int {
	// 总数
	var sum int = 0
	var left int = 0
	var centerIndex int = -1
	for _, v := range nums {
		sum += v
	}
	for i, _ := range nums {
		if i == 0 {
			left = 0
		} else {
			left += nums[i - 1]
		}
		right := sum - left - nums[i]
		if right == left {
			centerIndex = i
			break
		}
	}
	return centerIndex
}

func main()  {
	var nums []int
	nums = []int{-1,-1,0,0,-1,-1}
	var a map[string]string
	a = map[string]string{
		"1" : "1",
	}
	fmt.Println(a)
	fmt.Println(pivotIndex(nums))
}
