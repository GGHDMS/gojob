package main

import "fmt"
import "gojob/something"

var name string = "seokmun"

func main() {

	n := "sm"

	fmt.Println("Hello World")

	something.SayBye()
	something.SayHello()

	fmt.Println(name)
	fmt.Println(n)
}
