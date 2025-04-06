package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	s2, _ := reader.ReadString('\n')

	s1 = s1[:len(s1)-1]
	s2 = s2[:len(s2)-1]

	if len(s1) != len(s2) {
		fmt.Println("NO")
		return
	}

	if s1 == s2 {
		fmt.Println("NO")
		return
	}

	var diff []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}

	if len(diff) != 2 {
		fmt.Println("NO")
		return
	}

	pos1, pos2 := diff[0], diff[1]
	s1Chars := []byte(s1)
	s1Chars[pos1], s1Chars[pos2] = s1Chars[pos2], s1Chars[pos1]

	if string(s1Chars) == s2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
