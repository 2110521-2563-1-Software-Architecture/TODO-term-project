package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientReader := bufio.NewReader(conn)
	for {
		data, err := clientReader.ReadString('\n')
		switch err {
		case nil:
			data := strings.TrimSpace(data)
			if data == "STOP" {
				log.Println("Stopping...")
				return
			}
			log.Println(data)
		case io.EOF:
			log.Println("EOF")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		_, err = conn.Write([]byte("Response from tracker\n"))
		if err != nil {
			log.Printf("Failed to response: %v\n", err)
		}
	}
}
