package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64 = 4
	var t float64 = 20

	luasPermukaan := 2 * math.Pi * r * (r + t)

	fmt.Println(luasPermukaan)
}
