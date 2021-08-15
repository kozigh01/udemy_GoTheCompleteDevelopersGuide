package main

import (
	"fmt"
	// "io"
	"net/http"
	// "os"
)

func main() {
	fmt.Println("Welcome to concurrency")

	urls := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.orgs",
		"http://amazon.com",
	}

	for _, url := range urls {
		processUrl(url)
	}
	fmt.Println()
}

func processUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("There was an error calling %q (resp = %v): %v\n", url, resp, err)
		return
	}
	// fmt.Printf("Body of call to %q:\n", url)
	// io.Copy(os.Stdout, resp.Body)
	fmt.Printf("Call to %q resulted in response code %q\n", url, resp.Status)
}