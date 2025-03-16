package main

import (
	"fmt"
	"math"
)

func main() {
	// Read the number of sushi pieces
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	// Validate input range for n (2 ≤ n ≤ 100000)
	if !validateInput(n, 2, 100000) {
		return
	}

	// Read the type of each sushi piece (1 for tuna, 2 for eel)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		// Validate each sushi type is either 1 or 2
		if !validateInput(arr[i], 1, 2) {
			return
		}
	}

	// Count consecutive sequences of same type sushi
	sequences := make([]int, 0)
	currentCount := 1
	currentType := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] == currentType {
			currentCount++
		} else {
			// When type changes, store the count and reset
			sequences = append(sequences, currentCount)
			currentCount = 1
			currentType = arr[i]
		}
	}
	sequences = append(sequences, currentCount)

	// Find the maximum valid segment length
	// A valid segment has equal halves of different types
	// So we take minimum of adjacent sequences and multiply by 2
	maxLength := 0
	for i := 0; i < len(sequences)-1; i++ {
		length := int(math.Min(float64(sequences[i]), float64(sequences[i+1]))) * 2
		if length > maxLength {
			maxLength = length
		}
	}

	// Print the maximum length of valid segment
	fmt.Println(maxLength)
}

// validateInput checks if a value is within the specified range
func validateInput(value int, min int, max int) bool {
	return value >= min && value <= max
}
