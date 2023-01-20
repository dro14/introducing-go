package main

import "fmt"

func factorial() func() (int, int) {
	n, value := 0, 1
	return func() (int, int) {
		ret1, ret2 := n, value
		n++
		value *= n
		return ret1, ret2
	}
}

func main() {
	next := factorial()
	for i := 0; i <= 20; i++ {
		n, value := next()
		fmt.Printf("%d! = %d\n", n, value)
	}
}
