package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errHttp = errors.New("http erro")

func main() {

	result := map[string]string{}

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		code := "OK"

		err := hitUrl(url)

		if err != nil {
			code = "ERR"
		}

		result[url] = code
	}

	fmt.Println(result)

}

func hitUrl(url string) error {
	fmt.Println("Checking:", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errHttp
	}

	return nil
}
