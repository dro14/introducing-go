// In order to use the Exit function from the os package, you would need to import the os package: import "os" and the invoke it with .: os.Exit().

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Calling Exit() function from \"os\" package.")
	os.Exit(0)
}
