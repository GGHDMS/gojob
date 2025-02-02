package main

import "fmt"

func supperAdd(numbers ...int) int {
	var total int
	for i, n := range numbers {
		fmt.Println(i, n)
		total += n
	}

	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
