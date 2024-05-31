package main

import (
	"fmt"
)

func getMinMax(a1, a2, a3, a4, a5, a6 int) (min int, max int) {
	numbers := []int{a1, a2, a3, a4, a5, a6}
	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6 int

	fmt.Println("Enter 6 numbers:")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max := getMinMax(a1, a2, a3, a4, a5, a6)

	fmt.Printf("%d is the maximum number\n", max)
	fmt.Printf("%d is the minimum number\n", min)
}
