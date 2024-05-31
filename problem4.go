package main

import (
	"fmt"
)

type Cipher interface {
	encode() string
	decode() string
}

type student struct {
	name string
}

func (s student) encode() string {
	encodedName := ""
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			encodedName += string((char-'a'+3)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			encodedName += string((char-'A'+3)%26 + 'A')
		} else {
			encodedName += string(char)
		}
	}
	return encodedName
}

func (s student) decode() string {
	decodedName := ""
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			decodedName += string((char-'a'-3+26)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			decodedName += string((char-'A'-3+26)%26 + 'A')
		} else {
			decodedName += string(char)
		}
	}
	return decodedName
}

func main() {
	var choice int
	var name string

	fmt.Println("[1] Encrypt")
	fmt.Println("[2] Decrypt")
	fmt.Print("Choose your menu? ")
	fmt.Scan(&choice)

	fmt.Print("Input Student's Name: ")
	fmt.Scan(&name)

	s := student{name: name}

	switch choice {
	case 1:
		fmt.Printf("Encode of Student's Name %s is %s\n", name, s.encode())
	case 2:
		fmt.Printf("Decode of Student's Name %s is %s\n", name, s.decode())
	default:
		fmt.Println("Invalid choice")
	}
}
