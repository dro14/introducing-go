package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	if input < 0 {
		panic("input is less than zero")
	}

	arr := make([]int, input)
	for i := 0; i < input; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)
}
