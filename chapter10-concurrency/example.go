package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int64
}

func main() {
	urls := []string{
		"https://www.apple.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.microsoft.com",
		"https://www.facebook.com",
		"https://www.akfauniversity.uz",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(res.Body)

			var buf bytes.Buffer
			size, err := buf.ReadFrom(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: size,
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Printf("The biggest home page: %s with a size of %d bytes\n", biggest.URL, biggest.Size)
}
