package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
	}

	colors["yellow"] = "#FFFF00"

	fmt.Println(colors)

	someMap := make(map[int]string)
	someMap[1] = "One"
	someMap[2] = "Two"
	someMap[3] = "Three"

	delete(someMap, 2)

	fmt.Println(someMap)

	printKeyValues(colors)

}

func printKeyValues(m map[string]string) {
	for key, value := range m {
		fmt.Println("Value for", key, "is:", value)
	}
}
