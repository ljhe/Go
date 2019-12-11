package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main()  {
	haystack := "a"
	needle := "a"
	fmt.Print(strStr(haystack, needle))
}
