package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	x := a[:]
	fmt.Println(x)
	x = append(x, 6)
	fmt.Println(x)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	x = append(x, slice2...)
	fmt.Println(x)

	slice1 = []int{1, 2, 3}
	slice2 = make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	copy(slice1, x[5:8])
	fmt.Println(slice1, slice2)
}
