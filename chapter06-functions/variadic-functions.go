package main

import "fmt"

func add(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	//arr := []int{1,2,3,4,5}
	fmt.Println(add(1, 2, 3, 4, 5))
}
