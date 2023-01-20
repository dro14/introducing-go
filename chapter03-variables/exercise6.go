package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scan(&input)

	output := input * 0.3048

	fmt.Println(output)
}
