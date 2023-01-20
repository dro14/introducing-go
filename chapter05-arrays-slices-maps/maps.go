package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["first"] = 10
	x["second"] = 51
	x["third"] = 93
	x["fourth"] = 72
	x["fifth"] = 14
	value, ok := x["fifth"]
	fmt.Println(value, ok)

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	element, ok := elements["Li"]
	fmt.Println(element, ok)
}
