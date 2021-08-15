package interfaces

import (
	"fmt"
	"io"
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

	// var buffer = make([]byte, 99999)
	// n, _ := resp.Body.Read(buffer)
	// fmt.Printf("\nResponse: n=%v (len:%v, cap:%v)\n   body: %q\n", n, len(buffer), cap(buffer), string(buffer[:n]))

	// using built in go helpers to read data
	fmt.Println("\n---------------------------------------------")
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}

func LogWriterExample() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("There was an error calling get: %v\n", err)
		os.Exit(1)
	}

	lw := LogWriter{}

	io.Copy(&lw, resp.Body)
}