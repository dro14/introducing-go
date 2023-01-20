// Program that prints numbers 1 through 100, but replaces multiples of 3 with "Fizz", multiples of 5 with "Buzz", and multiples of both 3 and 5 with "FizzBuzz":

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		printed := false
		if i%3 == 0 {
			fmt.Print("Fizz")
			printed = true
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
			printed = true
		}
		if !printed {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
