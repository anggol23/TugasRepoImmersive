package main

import (
	"fmt"
)

func ArrayUnique(arrayA, arrayB []int) []int {
	mapB := make(map[int]bool)
	for _, val := range arrayB {
		mapB[val] = true
	}

	var result []int
	for _, val := range arrayA {
		if !mapB[val] {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 50}))
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))
}
