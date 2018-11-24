package main

import (
	"io"
	"log"
	"os"
)

func main() {
	a := os.Args
	if len(a) < 2 {
		log.Fatalln("Please specify a filename")
	}
	file, err := os.Open(a[1])
	if err != nil {
		log.Fatalf("Could not open file %v.\nGot following error: %v\n", a[1], err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
