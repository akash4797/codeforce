package main

import (
	"fmt"
)

func main() {
	// Read the happiness matrix
	g := make([][]int, 5)
	for i := range g {
		g[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			fmt.Scan(&g[i][j])
		}
	}

	// Generate all possible permutations and find maximum happiness
	maxHappiness := 0
	perm := []int{0, 1, 2, 3, 4}
	generatePermutations(perm, 0, &maxHappiness, g)

	fmt.Println(maxHappiness)
}

func generatePermutations(arr []int, start int, maxHappiness *int, g [][]int) {
	if start == len(arr) {
		happiness := calculateHappiness(arr, g)
		if happiness > *maxHappiness {
			*maxHappiness = happiness
		}
		return
	}

	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start]
		generatePermutations(arr, start+1, maxHappiness, g)
		arr[start], arr[i] = arr[i], arr[start]
	}
}

func calculateHappiness(order []int, g [][]int) int {
	total := 0

	// Initial line conversations
	for i := 0; i < 4; i += 2 {
		if i+1 < 5 {
			p1, p2 := order[i], order[i+1]
			total += g[p1][p2] + g[p2][p1]
		}
	}

	// After first person enters shower (4 people left)
	for i := 0; i < 3; i += 2 {
		if i+1 < 4 {
			p1, p2 := order[i+1], order[i+2]
			total += g[p1][p2] + g[p2][p1]
		}
	}

	// After second person enters shower (3 people left)
	for i := 0; i < 2; i += 2 {
		if i+1 < 3 {
			p1, p2 := order[i+2], order[i+3]
			total += g[p1][p2] + g[p2][p1]
		}
	}

	// After third person enters shower (2 people left)
	p1, p2 := order[3], order[4]
	total += g[p1][p2] + g[p2][p1]

	return total
}
