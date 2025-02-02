package main

import "fmt"

func main3() {
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

	m := map[string]int{"a": 1, "b": 2}

	for k, v := range m {
		fmt.Println(k, v)
	}

	p := Person{"seokmun", 21, []string{"korean", "english"}}
	fmt.Println(p)
	p = Person{name: "seokmun", age: 20, favFood: []string{"korean", "english"}}
	fmt.Println(p)
}

type Person struct {
	name    string
	age     int
	favFood []string
}
