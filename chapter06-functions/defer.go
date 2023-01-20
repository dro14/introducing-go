package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	start := time.Now()
	defer func() { fmt.Println("Time elapsed:", time.Since(start)) }()

	for i := 0; i < input; i++ {
		fmt.Println(factorial(input))
	}

	defer func() { fmt.Println("Deferred first, called third") }()
	defer func() { fmt.Println("Deferred second, called second") }()
	defer func() { fmt.Println("Deferred third, called first") }()
}
