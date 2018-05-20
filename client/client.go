package main

import (
	"log"

	"github.com/bsandusky/protocache/client/cache"
	"github.com/bsandusky/protocache/client/cli"
	"github.com/bsandusky/protocache/client/explorer"
)

func main() {

	err := cache.InitClient("localhost", "8080")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	go cli.Start(done)
	go explorer.Start()

	// Wait for exit command to close
	<-done
}
