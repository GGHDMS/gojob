package main

import (
	"fmt"
	"strings"
)
import "gojob/something"

var name = "seokmun"

func multifly(a, b int) int {
	return a + b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {

	n := "sm"

	fmt.Println("Hello World")

	something.SayBye()
	something.SayHello()

	fmt.Println(name)
	fmt.Println(n)

	fmt.Println(multifly(1, 2))
	fmt.Println(lenAndUpper(name))

	repeatMe("a", "b", "c")
}
