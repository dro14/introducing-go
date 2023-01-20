package main

import "fmt"

func factorial(n int, sending chan<- int) {
	if n == 0 {
		sending <- 1
	}
	receiving := make(chan int)
	go factorial(n-1, receiving)
	previousValue := <-receiving
	sending <- n * previousValue
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	receiving := make(chan int)
	go factorial(n, receiving)
	fmt.Println(<-receiving)
}
