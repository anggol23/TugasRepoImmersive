package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20

	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)

	swap(&a, &b)

	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
}
