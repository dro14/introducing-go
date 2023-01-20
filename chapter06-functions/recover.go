package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	defer func() {
		if x := recover(); x != nil {
			input = -1 * x.(int)
		}

		arr := make([]int, input)
		for i := 0; i < input; i++ {
			arr[i] = i + 1
		}
		fmt.Println(arr)
	}()
	if input < 0 {
		panic(input)
	}
}
