package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Expected a filename as the only command line argument")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
