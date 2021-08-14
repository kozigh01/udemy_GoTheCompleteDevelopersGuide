package interfaces

import (
	"fmt"
	"net/http"
	"os"
)

func HttpGet() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("There was an error calling get: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response:\n   status=%v\n   resp=%v\n", resp.Status, resp)
}