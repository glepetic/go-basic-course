package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// Channel is used to communicate between go routines (main and child ones)
	c := make(chan string)

	for _, url := range urls {
		// With go keyword, a go routine is created to execute the function asynchronously in another thread
		go checkStatus(url, c)
	}

	for u := range c {
		// Function Literal = Lambda in other languages
		// Add the innerUrl parameter to make a copy of u and guarantee that there is not a shared variable with a race condition
		go func(innerUrl string) {
			time.Sleep(5 * time.Second)
			checkStatus(innerUrl, c)
		}(u)
	}

}

// By default, this function is Sequential, Serial, Synchronous and order preserved
func checkStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Domain %v might be down!\n", url)
		c <- url
		return
	}
	fmt.Printf("Domain %v is up!\n", url)
	c <- url
}
