package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(bs)
	fmt.Println("just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
