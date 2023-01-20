package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			println(err)
			continue
		}
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		println(err)
	} else {
		println("Received", msg)
	}
	c.Close()
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		println(err)
		return
	}

	msg := "Hello, World"
	println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		println(err)
	}

	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scan(&input)
}
