package links

import (
	"fmt"
	"net/http"
	"time"
)

func Links1() {
	fmt.Println("Welcome to concurrency")

	startTime := time.Now()

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
	fmt.Printf("Time to finish: %.3f\n", time.Since(startTime).Seconds())
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