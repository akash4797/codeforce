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

	letterCount := make(map[rune]int)
	for _, char := range s1 {
		letterCount[char]++
	}

	possible := true
	for _, char := range s2 {
		if char == ' ' {
			continue
		}
		if letterCount[char] > 0 {
			letterCount[char]--
		} else {
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
