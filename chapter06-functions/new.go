package main

import "fmt"

func one(arr *[20]int) {
	for i := range arr {
		arr[i] = 1
	}
}

func array() *[20]int {
	arr := new([20]int)
	for i := range arr {
		arr[i] = i + 1
	}
	return arr
}

func main() {

	xPtr := array()
	fmt.Println(*xPtr)
	one(xPtr)
	fmt.Println(*xPtr)
}
