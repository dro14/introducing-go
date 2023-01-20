package main

import (
	"fmt"
	"time"
)

func Sleep(seconds float64) {
	<-time.After(time.Duration(seconds * float64(time.Second)))
}

func main() {
	start := time.Now()
	Sleep(4.3)
	elapsed := time.Since(start)
	fmt.Println("elapsed:", elapsed)
}
