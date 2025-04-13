package main

import (
	"fmt"
)

func main() {
	var binary string
	fmt.Scan(&binary)

	zeroIndex := -1
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			zeroIndex = i
			break
		}
	}

	if zeroIndex != -1 {
		result := binary[:zeroIndex] + binary[zeroIndex+1:]
		for len(result) > 1 && result[0] == '0' {
			result = result[1:]
		}
		fmt.Println(result)
	} else {
		result := binary[1:]
		for len(result) > 1 && result[0] == '0' {
			result = result[1:]
		}
		fmt.Println(result)
	}
}
