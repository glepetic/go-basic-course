package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args

	if len(args) == 1 || len(args) > 2 {
		fmt.Printf("Error: Invalid amount of arguments (%v). Expected total arguments: 1.\n", len(args)-1)
		os.Exit(1)
	}

	fileName := args[1]

	fp, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	io.Copy(os.Stdout, fp)

}
