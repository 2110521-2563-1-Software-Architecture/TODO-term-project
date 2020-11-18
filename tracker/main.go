package main

import (
	"context"
	"log"
	"net"
)

var ctx = context.Background()

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	log.Println("Listening...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
