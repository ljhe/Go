package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	left := 0
	right := length - 1
	var index []int
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			index = append(index, left + 1)
			index = append(index, right + 1)
			return index
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return index
}

func main()  {
	var numbers []int
	numbers = []int{2, 7, 11, 15}
	fmt.Print(twoSum(numbers, 9))
}
