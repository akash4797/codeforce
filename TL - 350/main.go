package main

import (
	"fmt"
)

func max(slice []int) int {
	maxVal := slice[0]
	for _, v := range slice {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func min(slice []int) int {
	minVal := slice[0]
	for _, v := range slice {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	correct := make([]int, n)
	wrong := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&correct[i])
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&wrong[i])
	}

	maxCorrect := max(correct)
	minCorrect := min(correct)
	minWrong := min(wrong)

	v := maxInt(maxCorrect, 2*minCorrect)

	if v < minWrong {
		fmt.Println(v)
	} else {
		fmt.Println(-1)
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
