package main

import "fmt"

func main() {
	a := 4
	b := &a

	*b = 5
	fmt.Println(*b)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	arr := [5]string{"a", "b", "c"}
	fmt.Println(arr)

	slice = append(slice, 4)
	fmt.Println(slice)
}
