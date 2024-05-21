package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n (1-30): ")
	fmt.Scanln(&n)

	if n < 1 || n > 30 {
		fmt.Println("Input tidak valid. Harap masukkan nilai n antara 1 dan 30.")
		return
	}

	fmt.Print("   ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	fmt.Print("  +")
	for i := 1; i <= n; i++ {
		fmt.Print("----")
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%2d|", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
