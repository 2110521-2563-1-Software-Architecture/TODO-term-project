package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type ResponseBody struct {
	Message string
}

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, &ResponseBody{Message: "Hello World"})
	// })
	// e.Logger.Fatal(e.Start(":9000"))

	go fileHandler();

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	log.Println("Running")
	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		request, err := clientReader.ReadString('\n')
		switch err {
		case nil:
			request := strings.TrimSpace(request)
			if _, err = conn.Write([]byte(request + "\n")); err != nil {
				log.Printf("failed to send: %v\n", err)
			}
		case io.EOF:
			log.Println("closing...")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		response, err := serverReader.ReadString('\n')
		switch err {
		case nil:
			log.Println(strings.TrimSpace(response))
		case io.EOF:
			log.Println("server closed")
			return
		default:
			log.Printf("server error: %v\n", err)
			return
		}
	}
}
