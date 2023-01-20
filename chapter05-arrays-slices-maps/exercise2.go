// The length of the slice is 3.

package main

import "fmt"

func main() {
	slice := make([]int, 3, 9)
	fmt.Println(len(slice))
}
