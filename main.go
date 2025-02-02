package main

import "fmt"

func main() {
	a := 4
	b := &a

	*b = 5
	fmt.Println(*b)
}
