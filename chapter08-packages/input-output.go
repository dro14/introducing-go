package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const s string = `These days are unusually cold. Even water on the windows has frozen.
I should have started my internship, but due to this cold, it was
posponed to warmer days.`

func main() {
	var x byte
	fmt.Scan(&x)

	r := strings.NewReader(s)

	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}
	buf.WriteByte(x)

	m, err := io.Copy(os.Stdout, buf)
	if err != nil {
		panic(err)
	}

	buf.Reset()
	_, err = buf.WriteString(strconv.Itoa(int(n)))
	if err != nil {
		panic(err)
	}
	buf.WriteByte(x)
	_, err = buf.WriteString(strconv.Itoa(int(m)))
	if err != nil {
		panic(err)
	}
	buf.WriteByte(x)

	_, err = io.Copy(os.Stdout, buf)
	if err != nil {
		panic(err)
	}
}
