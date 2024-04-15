package main

import "fmt"

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		age:       30,
		contactInfo: contactInfo{
			email: "sam.smith@go.com",
			zip:   1400,
		},
	}

	jim.introduce()
	fmt.Println(jim)
	jim.print()

	alex := jim.updateNameByValue("Alex")
	jim.introduce()
	alex.introduce()

	// alternative (&alex).updateNameByReference("Sam")
	alex.updateNameByReference("Sam") // Go allows shortcut for value to pointer conversion to use functions
	alex.introduce()                  // now Sam

}
