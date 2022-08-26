package main

import (
	"log"

	"github.com/gsxhnd/dgraph-example/example"
)

func main() {
	log.Println("Starting dgraph example...")
	c := example.NewClient()
	c.SimpleQuery()
}
