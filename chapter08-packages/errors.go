package main

import (
	"errors"
	"fmt"
)

func main() {
	str := fmt.Sprintf("error message %d", 1)
	err := errors.New(str)
	panic(err)
}
