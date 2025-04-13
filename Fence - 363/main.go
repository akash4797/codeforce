package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}

	sum := 0
	for i := 0; i < k; i++ {
		sum += h[i]
	}

	minSum := sum
	minIndex := 1

	for i := k; i < n; i++ {
		sum = sum + h[i] - h[i-k]
		if sum < minSum {
			minSum = sum
			minIndex = i - k + 2
		}
	}

	fmt.Fprintln(writer, minIndex)
}
