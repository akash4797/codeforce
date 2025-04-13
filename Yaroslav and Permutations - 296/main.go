package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	freq := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		freq[arr[i]]++
	}

	possible := true
	for _, count := range freq {
		if count > (n+1)/2 {
			possible = false
			break
		}
	}

	if possible {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
