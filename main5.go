package main

import (
	"fmt"
	"gojob/mydict"
)

func main5() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)

	search, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(search)
	}

	search, err = dictionary.Search("world")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(search)
	}

	fmt.Println("=================")
	err = dictionary.Add("hello", "hello")
	fmt.Println(err)
	err = dictionary.Add("world", "hello")
	fmt.Println(err)

	err = dictionary.Update("hello", "bye")
	search, err = dictionary.Search("hello")

	fmt.Println(search)

	err = dictionary.Update("no", "bye")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary.Search("no"))
	}

	dictionary.Delete("world")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary.Search("world"))
	}
}
