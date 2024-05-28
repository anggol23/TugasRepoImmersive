package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	shifted := make([]rune, len(input))
	offset = offset % 26

	for i, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted[i] = 'a' + (char-'a'+rune(offset))%26
		} else {
			shifted[i] = char
		}
	}

	return string(shifted)
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "siterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
