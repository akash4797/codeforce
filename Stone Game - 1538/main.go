package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		// Read the array of stone powers
		stones := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&stones[j])
		}

		// Find min and max values and their positions
		minVal, maxVal := stones[0], stones[0]
		minPos, maxPos := 0, 0
		for j := 0; j < n; j++ {
			if stones[j] < minVal {
				minVal = stones[j]
				minPos = j
			}
			if stones[j] > maxVal {
				maxVal = stones[j]
				maxPos = j
			}
		}

		// Calculate minimum moves needed
		result := calculateMinMoves(n, minPos, maxPos)
		fmt.Println(result)
	}
}

func calculateMinMoves(n, minPos, maxPos int) int {
	// Case 1: Take both from left side
	moves1 := max(minPos, maxPos) + 1

	// Case 2: Take both from right side
	moves2 := n - min(minPos, maxPos)

	// Case 3: Take one from left and one from right
	leftRightMoves := min(maxPos, minPos) + 1 + (n - max(maxPos, minPos))

	// Return the minimum of all cases
	return min(moves1, min(moves2, leftRightMoves))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
