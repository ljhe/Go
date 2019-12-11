package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	length := len(nums)
	sum := 0
	j := 0
	for i := 0; i <= length; i++ {
		if i != length && nums[i] == 1 {
			j++
		} else {
			if sum < j {
				sum = j
			}
			j = 0
		}
	}
	return sum
}

func main()  {
	nums := []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
