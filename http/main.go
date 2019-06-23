package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct{}

func (logWriter) Write(p []byte) (n int, err error) {
	fmt.Println(p)

	return len(p), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
