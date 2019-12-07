package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	for n % 4 == 0 {
		n /= 4
	}
	if n % 8 == 7 {
		return 4
	}
	a := 0
	for a * a <= n {
		b := int(math.Sqrt(float64(n - a * a)))
		if a * a + b * b == n {
			m, n := 0, 0
			if a != 0 {
				m = 1
			} else {
				m = 0
			}
			if b != 0 {
				n = 1
			} else {
				n = 0
			}
			return m + n
		}
		a += 1
	}
	return 3
}

func main()  {
	fmt.Println(numSquares(13))
}
