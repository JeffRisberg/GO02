package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a filename as first argument.")
		os.Exit(1)
	}
	filename := os.Args[1]

	fmt.Println(filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(file)

	io.Copy(os.Stdout, file)
}
