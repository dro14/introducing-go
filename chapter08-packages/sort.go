package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	size, _ := strconv.Atoi(os.Args[1])
	arr := make([]int, size)

	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		arr[i] = rand.Intn(size)
	}

	//fmt.Println(arr)
	start := time.Now()
	sort.Ints(arr)
	elapsed := time.Since(start).Seconds()
	//fmt.Println(arr)
	fmt.Printf("Merge sort took %.6f seconds to sort a slice of size %d with random values.\n", elapsed, size)
}
