package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	//var c1 Circle
	//c2 := new(Circle)
	//c3 := Circle{x: 0, y: 0, r: 5}
	//c4 := Circle{0, 0, 5}
	c5 := &Circle{0, 0, 5}

	fmt.Println(c5.x, c5.y, c5.r)
	c5.x = 10
	c5.y = 5
	fmt.Println(c5.x, c5.y, c5.r)
	fmt.Println(circleArea(c5))
}
