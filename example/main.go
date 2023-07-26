package main

import (
	"log"
	"time"

	"github.com/dillonstreator/fnopt/example/api"
)

func main() {
	_, err := api.NewServer(
		":3000",
		api.ServerWithMaxConns(25),
		api.ServerWithTimeout(time.Second*30),
	)
	if err != nil {
		log.Fatal(err)
	}
}
