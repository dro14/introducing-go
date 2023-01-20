// The value of the expression is true.

package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}
