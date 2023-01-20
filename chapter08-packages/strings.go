package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Contains(s, substr string) bool
	fmt.Println(strings.Contains("MacOS Ventura", "Mac"))
	// => true
	fmt.Println(strings.Contains("MacOS Ventura", "W"))
	// => false

	// func Count(s, sep string) int
	fmt.Println(strings.Count("Mississippi", "i"))
	// => 4
	fmt.Println(strings.Count("Mississippi", "a"))
	// => 0

	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("unbelievable", "un"))
	// => true
	fmt.Println(strings.HasPrefix("unbelievable", "in"))
	// => false

	// func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("probable", "able"))
	// => true
	fmt.Println(strings.HasSuffix("probable", "ous"))
	// => false

	// func Index(s, sep string) int
	fmt.Println(strings.Index("disambiguation", "tion"))
	// => 10
	fmt.Println(strings.Index("disambiguation", "anti"))
	// => -1

	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"D", "o", "n", "i", "y", "o", "r", "b", "e", "k"}, ""))
	// => Doniyorbek
	fmt.Println(strings.Join([]string{"12", "01", "2023"}, "."))
	// => 12.01.2023

	// func Repeat(s string, count int) string
	fmt.Println(strings.Repeat("yo", 2))
	// => yoyo
	fmt.Println(strings.Repeat("mur", 2))
	// => murmur

	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("chapter05-arrays-slices-maps", "-", " ", -1))
	// => chapter05 arrays slices maps
	fmt.Println(strings.Replace("2023-1-12", "-", ".", 1))
	// => 2023.1-12

	// func Split(s, sep string) []string
	fmt.Println(strings.Split("1/12/2023", "/"))
	// => [1 12 2023]
	fmt.Println(strings.Split("AKFA University", ""))
	// => [A K F A   U n i v e r s i t y]

	// func ToLower(s string) string
	fmt.Println(strings.ToLower("ALL THIS TEXT SHOULD BE IN LOWER CASE: ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	// => all this text should be in lower case: abcdefghijklmnopqrstuvwxyz

	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("all this text should be in upper case: abcdefghijklmnopqrstuvwxyz"))
	// => ALl THIS TEXT SHOULD BE IN UPPER CASE: ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Convert a string to a slice of bytes
	str := "Doniyorbek"
	fmt.Println([]byte(str))
	// => [68 111 110 105 121 111 114 98 101 107]

	// Convert a slice of bytes to a string
	arr := []byte{'R', 'a', 'k', 'h', 'm', 'o', 'n', 'b', 'e', 'r', 'd', 'i', 'e', 'v'}
	fmt.Println(string(arr))
	// => Rakhmonberdiev
}
