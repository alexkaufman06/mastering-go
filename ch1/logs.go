package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello world!")
	}
	log.Panic("Panic: Hello World!") // panic includes additional low-level information for debugging
}