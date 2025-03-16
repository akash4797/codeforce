package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Create a map to track which levels can be passed
	canPass := make(map[int]bool)

	// Process Little X's levels
	var p int
	fmt.Scan(&p)
	for i := 0; i < p; i++ {
		var level int
		fmt.Scan(&level)
		canPass[level] = true
	}

	// Process Little Y's levels
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var level int
		fmt.Scan(&level)
		canPass[level] = true
	}

	// Check if all levels can be passed
	canPassAll := true
	for i := 1; i <= n; i++ {
		if !canPass[i] {
			canPassAll = false
			break
		}
	}

	if canPassAll {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}
