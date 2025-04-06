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

	arr := make([]int, n)
	count := make([]int, 8)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
		count[arr[i]]++
	}

	groups := [][3]int{{1, 2, 4}, {1, 2, 6}, {1, 3, 6}}
	result := make([][3]int, 0, n/3)

	if count[5] > 0 || count[7] > 0 {
		fmt.Println(-1)
		return
	}

	for count[1] > 0 {
		found := false
		for _, group := range groups {
			if count[group[0]] > 0 && count[group[1]] > 0 && count[group[2]] > 0 {
				count[group[0]]--
				count[group[1]]--
				count[group[2]]--
				result = append(result, group)
				found = true
				break
			}
		}
		if !found {
			fmt.Println(-1)
			return
		}
	}

	for i := 1; i <= 7; i++ {
		if count[i] != 0 {
			fmt.Println(-1)
			return
		}
	}

	for _, group := range result {
		fmt.Printf("%d %d %d\n", group[0], group[1], group[2])
	}
}
