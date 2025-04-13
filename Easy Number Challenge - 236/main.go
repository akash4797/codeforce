package main

import (
	"fmt"
)

func countDivisors(n int, memo map[int]int) int {
	if val, exists := memo[n]; exists {
		return val
	}

	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count += 2
			}
		}
	}

	memo[n] = count
	return count
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	sum := 0
	mod := 1073741824
	memo := make(map[int]int)

	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			for k := 1; k <= c; k++ {
				product := i * j * k
				divisors := countDivisors(product, memo)
				sum = (sum + divisors) % mod
			}
		}
	}

	fmt.Println(sum)
}
