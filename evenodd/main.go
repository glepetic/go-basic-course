package main

import "fmt"

func main() {

	nums := []int{24, 16, 23, 94, 1245, 12, 13, 55, 626, 123, 1}

	for i, number := range nums {
		var result string
		if num(number).isEven() {
			result = "even"
		} else {
			result = "odd"
		}
		fmt.Println("Index:", i, "Number:", number, "Result:", result)
	}

}
