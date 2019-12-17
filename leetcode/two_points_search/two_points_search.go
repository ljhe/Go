package main

import "math"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := int(math.Floor(float64((left + right) / 2)))
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

func main()  {
	
}
