package main

import (
	m "doniyorbek.me/math"
	"fmt"
)

const BytesInKB = 1024

func PrintBytesInKB(size int) {
	fmt.Println(size * BytesInKB)
}

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
	min := m.Min(xs)
	fmt.Println(min)
	max := m.Max(xs)
	fmt.Println(max)
}
