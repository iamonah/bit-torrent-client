package main

import (
	"log"
	"os"

	"github.com/onahvictor/torrent-client/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
	err = tf.DownlaodFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
