package main

import "fmt"

func swap(x, y byte) (byte, byte) {
	return y, x
}

func reverseString(s []byte)  {
	length := len(s)
	left := 0
	right := length - 1
	for left < right {
		s[left], s[right] = swap(s[left], s[right])
		left++
		right--
	}
}

func main()  {
	var s []byte
	s = []byte{'h','e','l','l','o'}
	reverseString(s)
	fmt.Printf("%s",s)
}
