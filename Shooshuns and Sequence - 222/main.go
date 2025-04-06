package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	target := arr[k-1]

	for i := k - 1; i < n; i++ {
		if arr[i] != target {
			fmt.Println(-1)
			return
		}
	}

	operations := 0
	for i := k - 2; i >= 0; i-- {
		if arr[i] != target {
			operations = i + 1
			break
		}
	}

	fmt.Println(operations)
}
