package main

import (
	"fmt"
)

func pangkat(x, n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	base := x
	for n > 0 {
		if n%2 == 1 {
			result *= base
		}
		base *= base
		n /= 2
	}
	return result
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}
