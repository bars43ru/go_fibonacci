package main

import (
	"fmt"
)

func main() {
	//[1, 1, 2, 3, 5, 8, 13, 21 ]
	// 1  2  3  4  5  6   7   8
	fmt.Print(fibonacci(7))
}

func fibonacci(index int) int {
	if index < 1 {
		return 0
	} else if index <= 2 {
		return 1
	}
	v1 := fibonacci(index - 1)
	v2 := fibonacci(index - 2)
	return v1 + v2
}

