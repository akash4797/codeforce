package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var ticket string
	fmt.Scan(&n)
	fmt.Scan(&ticket)

	firstHalf := make([]int, n)
	secondHalf := make([]int, n)

	for i := 0; i < n; i++ {
		firstHalf[i] = int(ticket[i] - '0')
		secondHalf[i] = int(ticket[i+n] - '0')
	}

	sort.Ints(firstHalf)
	sort.Ints(secondHalf)

	allLess := true
	allGreater := true

	for i := 0; i < n; i++ {
		if firstHalf[i] >= secondHalf[i] {
			allLess = false
		}
		if firstHalf[i] <= secondHalf[i] {
			allGreater = false
		}
	}

	if allLess || allGreater {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
