package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	contests := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &contests[i])
	}

	// Sort contests in ascending order
	sort.Ints(contests)

	days := 0
	currentDay := 1

	// For each contest, check if we can use it for the current day
	for i := 0; i < n; i++ {
		if contests[i] >= currentDay {
			days++
			currentDay++
		}
	}

	fmt.Fprintln(writer, days)
}
