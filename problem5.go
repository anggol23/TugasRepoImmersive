package main

import (
	"fmt"
)

func MeanMedian(arrayInput []float64) (float64, float64) {
	n := len(arrayInput)
	if n == 0 {
		return 0, 0
	}

	sum := 0.0
	for _, num := range arrayInput {
		sum += num
	}
	mean := sum / float64(n)

	var median float64
	if n%2 == 0 {
		median = (arrayInput[n/2-1] + arrayInput[n/2]) / 2
	} else {
		median = arrayInput[n/2]
	}

	return mean, median
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))
	fmt.Println(MeanMedian([]float64{15, 20, 30, 68, 120}))
}
