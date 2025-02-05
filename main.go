package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)

	names := []string{"seokmun", "jaeho"}

	for _, n := range names {
		go isSexy(n, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)

}

func isSexy(user string, c chan bool) {
	time.Sleep(time.Second * 5)

	fmt.Println(user)
	c <- true
}
