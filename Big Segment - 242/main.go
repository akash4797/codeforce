package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	left := make([]int, n)
	right := make([]int, n)

	minLeft := int(1e9 + 1)
	maxRight := 0

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &left[i], &right[i])
		if left[i] < minLeft {
			minLeft = left[i]
		}
		if right[i] > maxRight {
			maxRight = right[i]
		}
	}

	answer := -1
	for i := 0; i < n; i++ {
		if left[i] == minLeft && right[i] == maxRight {
			answer = i + 1
			break
		}
	}

	fmt.Println(answer)
}
