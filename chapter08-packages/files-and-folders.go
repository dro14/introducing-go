package main

import (
	"os"
	"path/filepath"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Open the file
	file, err := os.Open("new-file.go")
	handleError(err)
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	handleError(err)

	// Make a slice of bytes of the size of the file
	bs := make([]byte, stat.Size())

	// Read the file
	_, err = file.Read(bs)
	handleError(err)

	str := string(bs)
	println(str)

	// With one function open, read the file and copy it to a slice of bytes
	ds, err := os.ReadFile("new-file.go")
	handleError(err)

	str = string(ds)
	println(str)

	new_file, err := os.Create("new-file.go")
	handleError(err)
	defer new_file.Close()

	new_file.WriteString("package main; func main() {println(\"Hello, World\")}")

	dir, err := os.Open(".")
	handleError(err)
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	handleError(err)

	for _, v := range fileInfos {
		println(v.Name())
	}

	filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		println(path)
		return nil
	})
}
