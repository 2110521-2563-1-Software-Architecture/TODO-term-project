package services

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Port() string {
	p := os.Getenv("USER_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":9000"
	}
	return fmt.Sprintf(":%s", p)
}

func Hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
