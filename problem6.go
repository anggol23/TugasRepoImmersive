package main

import (
	"fmt"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FullPrima(n int) string {
	if !isPrime(n) {
		return "Tidak"
	}

	str := strconv.Itoa(n)
	for _, digit := range str {
		digitInt, _ := strconv.Atoi(string(digit))
		if !isPrime(digitInt) {
			return "Tidak"
		}
	}

	return "Ya"
}

func main() {
	fmt.Println(FullPrima(2))
	fmt.Println(FullPrima(3))
	fmt.Println(FullPrima(11))
	fmt.Println(FullPrima(13))
	fmt.Println(FullPrima(23))
	fmt.Println(FullPrima(29))
	fmt.Println(FullPrima(37))
	fmt.Println(FullPrima(41))
	fmt.Println(FullPrima(43))
	fmt.Println(FullPrima(53))
}
