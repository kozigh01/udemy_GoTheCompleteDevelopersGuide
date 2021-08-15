package links

import (
	"fmt"
	"net/http"
	"time"
)

func Links3() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("Program execution duration: %.3f\n", time.Since(startTime).Seconds())
	}()

	called := make(chan string)
	// quit := time.NewTimer(3 * time.Second)

	urls := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.orgs",
		"http://amazon.com",
	}

	for _, url := range urls {
		go processUrl3(url, called)
	}

	for l := range called {		
		go func(url string) {
			time.Sleep(time.Duration(5) * time.Second)
			go processUrl3(url, called)
		}(l)
	}

	// for {
	// 	select {
	// 	case url := <-called:
	// 		time.Sleep(1 * time.Second)
	// 		go processUrl3(url, called)
	// 	case <- quit.C:
	// 		close(called)
	// 		fmt.Println("We are done.")
	// 	}
	// }
}

func processUrl3(url string, called chan<- string) {
	defer func () { called <- url }()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("There was an error calling %q (resp = %v): %v\n", url, resp, err)
		return
	}
	fmt.Printf("Call to %q resulted in response code %q\n", url, resp.Status)	
}