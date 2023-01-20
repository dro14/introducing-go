package main

import "fmt"

func main() {
	var x1 string = "Hello, World"
	fmt.Println(x1)

	var x2 string
	x2 = "Hello, World"
	fmt.Println(x2)

	var x3 string
	x3 = "first"
	fmt.Println(x3)
	x3 = "second"
	fmt.Println(x3)

	var x4 string
	x4 = "first "
	fmt.Println(x4)
	x4 = x4 + "second"
	fmt.Println(x4)

	var a string = "hello"
	var b string = "world"
	fmt.Println(a == b)
	b = "hello"
	fmt.Println(a == b)

	x := "Its type was inferred"
	fmt.Println(x)

	y := 5
	fmt.Println(y)
}
