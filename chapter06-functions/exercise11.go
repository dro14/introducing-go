// Program to swap two integers:

package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x, y := 1, 2
	fmt.Printf("x = %d; y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("x = %d; y = %d\n", x, y)
}
