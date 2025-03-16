package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	if !validateInput(t, 1, 10000) {
		return
	}

	totalN := 0
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		totalN += n

		if !validateInput(n, 1, 200000) || totalN > 200000 {
			return
		}

		processTestCase(n)
	}
}

func validateInput(value int, min int, max int) bool {
	return value >= min && value <= max
}

func processTestCase(n int) {
	// Create array to store elements
	arr := make([]int64, n)

	// Read n integers
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if !validateInput(int(arr[i]), -1000000000, 1000000000) {
			return
		}
	}

	maxSum, minOps := findMaxSumAndOps(arr)
	fmt.Println(maxSum, minOps)
}

func findMaxSumAndOps(arr []int64) (int64, int) {
	n := len(arr)
	maxSum := int64(0)
	ops := 0

	// For each element, we want its absolute value to be positive
	// to maximize the sum
	for i := 0; i < n; i++ {
		maxSum += abs(arr[i])
	}

	// Count segments of negative numbers
	// Each segment needs one operation to make all numbers positive
	isNegative := false
	for i := 0; i < n; i++ {
		if (arr[i] < 0 && !isNegative) || (arr[i] > 0 && isNegative) {
			if !isNegative {
				ops++
			}
			isNegative = !isNegative
		}
	}

	return maxSum, ops
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
