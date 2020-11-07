package main

import (
	"log"
	"os"

	"github.com/2110521-2563-1-Software-Architecture/TODO-term-project/tree/master/backend/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
