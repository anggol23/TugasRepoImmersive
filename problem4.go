package main

import (
	"fmt"
	"strings"
)

func createCipherMap() map[rune]rune {
	standardAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shiftedAlphabet := "KLMNOPQRSTUVWXYZABCDEFGHIJ"

	cipherMap := make(map[rune]rune)
	for i, char := range standardAlphabet {
		cipherMap[char] = rune(shiftedAlphabet[i])
	}
	return cipherMap
}

func ubahHuruf(message string) string {
	cipherMap := createCipherMap()
	var encryptedMessage strings.Builder
	for _, char := range message {
		if char == ' ' {
			encryptedMessage.WriteRune(' ')
		} else {
			encryptedMessage.WriteRune(cipherMap[char])
		}
	}
	return encryptedMessage.String()
}

func main() {
	fmt.Println(ubahHuruf("SEPULSA OKE"))
	fmt.Println(ubahHuruf("ALTERRA ACADEMY"))
	fmt.Println(ubahHuruf("INDONESIA"))
	fmt.Println(ubahHuruf("GOLANG"))
	fmt.Println(ubahHuruf("PROGRAMMER"))
}
