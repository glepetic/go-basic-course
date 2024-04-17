package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("https://google.com/")

	if err != nil {
		fmt.Println("Get request error:", err)
		os.Exit(1)
	}

	fmt.Println("Full Response:", res)
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	fmt.Println("Status:", res.Status)
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")

	/*
		--------------------
		Vanilla Read example
		--------------------

		bs := make([]byte, 99999)
		readBytes, err := res.Body.Read(bs)

		if err != nil && err != io.EOF {
			fmt.Println("Body read error:", err)
			os.Exit(2)
		}

		fmt.Println("Body read bytes:", readBytes)
		fmt.Println("Body:", string(bs))
	*/

	/*
		---------------------------
		Copy To StdOut Read Example
		---------------------------

		io.Copy(os.Stdout, res.Body)
	*/

	// Custom Writer Example

	lw := logWriter{}

	io.Copy(lw, res.Body)

}
