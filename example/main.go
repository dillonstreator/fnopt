package main

import (
	"log"
	"time"

	"github.com/dillonstreator/fnopt/example/somepkg"
)

func main() {
	_, err := somepkg.NewServer(
		":3000",
		somepkg.ServerWithMaxConns(25),
		somepkg.ServerWithTimeout(time.Second*30),
	)
	if err != nil {
		log.Fatal(err)
	}
}
