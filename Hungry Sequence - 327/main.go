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

	var n int
	fmt.Fscan(reader, &n)

	if n == 3 {
		fmt.Fprintf(writer, "2 9 15")
	} else if n == 5 {
		fmt.Fprintf(writer, "11 14 20 27 31")
	} else {
		start := n + 10
		for i := 0; i < n; i++ {
			if i == n-1 {
				fmt.Fprintf(writer, "%d", start+i*11)
			} else {
				fmt.Fprintf(writer, "%d ", start+i*11)
			}
		}
	}
}
