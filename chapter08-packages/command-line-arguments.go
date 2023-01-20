package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	size := flag.Int("size", 1, "the array size")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, *size)
	for i := range arr {
		arr[i] = rand.Intn(*size)
	}
	fmt.Println(arr)

	args := flag.Args()
	flag.Parse()
	fmt.Println(args)
}
