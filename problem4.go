package main

import (
	"fmt"
)

func findMaxSumSubArray(k int, arr []int) int {
	maxSum := 0
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < k; windowEnd++ {
		windowSum += arr[windowEnd]
	}
	maxSum = windowSum

	for windowEnd := k; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] - arr[windowStart]
		windowStart++
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2}))
	fmt.Println(findMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))
	fmt.Println(findMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))
	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))
	fmt.Println(findMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))
}
