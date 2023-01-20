// The sum function signature in Go looks like: func sum(xs []int) int.

package main

import (
	"fmt"
	"math/rand"
)

func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Int() % n
	}

	fmt.Println(sum(arr))
}
