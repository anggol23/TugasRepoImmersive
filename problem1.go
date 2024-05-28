package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}

	longestCommonSubstring := ""
	aLen := len(a)

	for length := aLen; length > 0; length-- {
		for i := 0; i <= aLen-length; i++ {
			substr := a[i : i+length]
			if strings.Contains(b, substr) {
				return substr
			}
		}
	}

	return longestCommonSubstring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
