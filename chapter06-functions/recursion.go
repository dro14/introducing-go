package main

import "fmt"

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findGCD(a int, b int) int {
	if b > a {
		return gcd(b, a)
	} else {
		return gcd(a, b)
	}
}

func main() {
	fmt.Println(findGCD(60, 48))
	fmt.Println(findGCD(24, 56))
}
