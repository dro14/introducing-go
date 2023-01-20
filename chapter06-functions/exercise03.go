// Function with a variadic parameter that finds the greatest number in a list of numbers:

package main

import "fmt"

func findMax(args ...int) int {
	max := args[0]
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(findMax(24, 62, 2, 236, 112, 63, 65, 72, 82, 14, 733, 26, 83, 13, 347, 13, 91))
}
