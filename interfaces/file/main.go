package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args[1]) // prints file name
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
	// go run main.go exfile.txt.

	// to print source code of main.go
	// go build main.go > ./main main.go
}
