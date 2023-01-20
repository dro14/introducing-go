// Defer defers a function to the moment before the surrounding function returns.
// Panic causes the function to immediately terminate along with any calling function
// and ultimately, the program itself. Recover is a way to prevent panics from going
// any further up the stack of functions. You can recover from a panic by deferring
// a call to a function which then calls the builtin recover function:

package main

import "fmt"

func f() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("RECOVERED:", x)
		}
	}()
	panic("panic!")
}

func main() {
	f()
}
