package main

import (
	"fmt"
)

func supperAdd(numbers ...int) int {
	var total int
	for i, n := range numbers {
		fmt.Println(i, n)
		total += n
	}

	return total
}

func main() {
	total := supperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
