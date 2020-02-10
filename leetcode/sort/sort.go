package main

import "fmt"

// 冒泡排序
func bubbleSort(arr []int)  {
	length := len(arr)
	// 临时变量
	var temp int
	// 一共需要比较几轮
	for i := 0; i < length - 1; i++ {
		// 每轮中比较找出最小值
		for j := 0; j < length - 1 - i ; j ++ {
			if arr[j] > arr[j + 1] {
				temp = arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}
	fmt.Println(arr)
}

// 选择排序
func selectSort(arr []int)  {
	length := len(arr)
	// 临时变量
	var temp int
	// 一共需要比较几轮
	for i := 0; i < length - 1; i++ {
		minIndex := i
		// 每一轮中找出最小的值
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		temp = arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp
	}
	fmt.Println(arr)
}

// 插入排序
func insertSort(arr []int)  {
	// 已经有序的数列 拿出一个新的数插入该数列中
	length := len(arr)
	for i := 1; i < length; i++ {
		tmp := arr[i] 	 // 从第二个数开始 从左向右依次取数
		key := i - 1     // 下标从0开始 依次从左向右
		// 每次取到的数都跟左侧的数做比较 如果左侧的数比取的数大 就将左侧的数右移一位 直至左侧没有数字比取的数大为止
		for key >= 0 && tmp < arr[key] {
			arr[key + 1] = arr[key]
			key--
		}
		// 将取到的数插入到不小于左侧数的位置
		arr[key + 1] = tmp
	}
	fmt.Println(arr)
}

func main()  {
	var arr []int
	arr = []int{1, 0, 3, 7, 6, 5, 2, 9, 4, 8}
	// 冒泡排序
	//bubbleSort(arr)
	// 选择排序
	//selectSort(arr)
	// 插入排序
	insertSort(arr)
}
