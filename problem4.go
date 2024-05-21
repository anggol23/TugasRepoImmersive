package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	filtered := ""
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			filtered += string(r)
		}
	}

	length := len(filtered)
	for i := 0; i < length/2; i++ {
		if filtered[i] != filtered[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(isPalindrome("civic"))
	fmt.Println(isPalindrome("katak"))
	fmt.Println(isPalindrome("kasur rusak"))
	fmt.Println(isPalindrome("kupu-kupu"))
	fmt.Println(isPalindrome("lion"))
}
