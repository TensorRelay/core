package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Printf("TensorRelay/core starting up\n")

	if err := run(); err != nil {
		log.Fatalf("fatal: %v", err)
		os.Exit(1)
	}
}

func run() error {
	// TODO: implement
	return nil
}
