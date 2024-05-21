package main

import (
	"fmt"
)

func drawXYZ(N int) {
	pattern := []rune{'Y', 'Z', 'X'}
	length := len(pattern)

	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			fmt.Print(string(pattern[(row+col)%length]))
		}
		fmt.Println()
	}
}

func main() {

	fmt.Println("Output untuk N = 3:")
	drawXYZ(3)

	fmt.Println("Output untuk N = 5:")
	drawXYZ(5)
}
