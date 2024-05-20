package main

import (
	"fmt"
	"sort"
)

func factors(n int) []int {
	var result []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
			if i != n/i {
				result = append(result, n/i)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}

func main() {
	var bilangan int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scanf("%d", &bilangan)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	factorsOfNumber := factors(bilangan)

	for _, factor := range factorsOfNumber {
		fmt.Println(factor)
	}
}

