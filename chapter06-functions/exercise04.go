// makeOddGenerator function that generates odd numbers.

package main

import "fmt"

func makeOddGenerator() func() int {
	i := -1
	return func() int {
		i += 2
		return i
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
