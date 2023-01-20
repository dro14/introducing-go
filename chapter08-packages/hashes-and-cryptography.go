package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getCrc32(fileName string) uint32 {
	file, err := os.Open(fileName)
	handleError(err)
	defer file.Close()

	hasher := crc32.NewIEEE()

	_, err = io.Copy(hasher, file)
	handleError(err)

	hash := hasher.Sum32()
	return hash
}

func getSha1(fileName string) []byte {
	file, err := os.Open(fileName)
	handleError(err)
	defer file.Close()

	hasher := sha1.New()

	_, err = io.Copy(hasher, file)
	handleError(err)

	hash := hasher.Sum([]byte{})
	return hash
}

func main() {
	var data string
	fmt.Scan(&data)

	hasher := crc32.NewIEEE()
	hasher.Write([]byte(data))
	v := hasher.Sum32()
	fmt.Println(v)

	var file1 string
	fmt.Scan(&file1)
	hash1 := getCrc32(file1)

	var file2 string
	fmt.Scan(&file2)
	hash2 := getCrc32(file2)

	fmt.Println(hash1, hash2, hash1 == hash2)

	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum([]byte{})
	fmt.Println(bs)

	h1 := getSha1(file1)
	h2 := getSha1(file2)

	areEqual := true
	for i := range h1 {
		if h1[i] != h2[i] {
			areEqual = false
			break
		}
	}

	fmt.Println(h1, h2, areEqual)
}
