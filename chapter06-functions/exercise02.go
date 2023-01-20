// Function that halves an integer and returns true if even value or false if odd value:

package main

import "fmt"

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(half(i))
	}
}
