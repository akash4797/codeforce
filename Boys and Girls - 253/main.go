package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var n, m int
	fmt.Sscanf(string(input), "%d %d", &n, &m)

	result := make([]byte, n+m)
	pos := 0

	if n > m {
		for m > 0 {
			result[pos] = 'B'
			pos++
			n--

			result[pos] = 'G'
			pos++
			m--
		}
		for n > 0 {
			result[pos] = 'B'
			pos++
			n--
		}
	} else {
		for n > 0 {
			result[pos] = 'G'
			pos++
			m--

			result[pos] = 'B'
			pos++
			n--
		}
		for m > 0 {
			result[pos] = 'G'
			pos++
			m--
		}
	}

	err = os.WriteFile("output.txt", []byte(string(result)+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
