package main

import "fmt"

// 剔除 slice 中指定的某一项
func removeElement(nums []int, val int) int {
	j := 0
	for _, v := range nums {
		if v != val {
			nums[j] = v
			j++
		}
	}
	return len(nums[:j])
}

func main()  {
	nums := []int{3,2,2,3}
	fmt.Print(removeElement(nums, 2))
}
