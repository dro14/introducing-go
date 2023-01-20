package main

import "fmt"

func main() {
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
