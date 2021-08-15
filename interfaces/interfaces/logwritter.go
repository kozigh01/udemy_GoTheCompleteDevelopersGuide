package interfaces

import "fmt"

type LogWriter struct {}

func (lw *LogWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Printf("Just wrote this many bytes: %v\n", len(b))
	return len(b), nil
}