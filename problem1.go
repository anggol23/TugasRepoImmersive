package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris segitiga: ")
	fmt.Scanln(&n)

	segitiga := make([][]string, n)

	for i := 0; i < n; i++ {
		segitiga[i] = make([]string, i+1)
		for j := 0; j < i+1; j++ {
			segitiga[i][j] = "*"
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(segitiga[i][j] + " ")
		}
		fmt.Println()
	}
}
