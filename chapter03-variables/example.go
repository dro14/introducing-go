package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scan(&input)

	output := input * 2

	fmt.Println(output)
}
