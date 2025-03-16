package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	// Binary search to find number of candies eaten
	left, right := 0, n
	for left <= right {
		eaten := (left + right) / 2
		puts := n - eaten // number of times candies were put

		// Sum of arithmetic sequence: first=1, diff=1, count=puts
		totalPut := (puts * (1 + puts)) / 2

		// If we put totalPut candies and ate 'eaten' candies
		remaining := totalPut - eaten

		if remaining == k {
			fmt.Println(eaten)
			return
		}

		if remaining > k {
			left = eaten + 1
		} else {
			right = eaten - 1
		}
	}
}
