package main

import "fmt"

// 存在返回 true 否则返回 false
func check(arr []int, v int) bool {
	var result bool
	for _, i := range arr {
		if i == v {
			result = true
			return result
		} else {
			result = false
		}
	}
	return result
}

func removeDuplicates(nums []int) int {
	j := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if j == 0 {
			nums[j] = nums[i]
			j++
		} else if !check(nums[:j], nums[i]) {
			nums[j] = nums[i]
			j++
		}
	}
	return len(nums[:j])
}

func main()  {
	nums := []int{0,0,1,1}
	fmt.Println(removeDuplicates(nums))
}
