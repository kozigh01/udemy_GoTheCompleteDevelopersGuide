package links

import (
	"fmt"
	"net/http"
	"time"
)

func Links2() {
	fmt.Println("Welcome to concurrency")

	startTime := time.Now()
	defer func() {
		fmt.Printf("Program execution duration: %.3f\n", time.Since(startTime).Seconds())
	}()

	timing := make(chan float64)

	urls := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, url := range urls {
		go processUrl2(url, timing)
	}

	var totalTime float64
	for i := 0; i < len(urls); i++ {
		totalTime += <-timing		
	}

	fmt.Printf("Execution duration for summarization of individual times: %.3f\n", totalTime)
}

func processUrl2(url string, timing chan<- float64) {
	startTime := time.Now()
	defer func () { timing <- time.Since(startTime).Seconds() }()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("There was an error calling %q (resp = %v): %v\n", url, resp, err)
		return
	}
	fmt.Printf("Call to %q resulted in response code %q\n", url, resp.Status)	
}