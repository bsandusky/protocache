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
	go cli.Run(done)
	go explorer.Run(done)

	// Wait for exit commands to close
	for i := 0; i < 2; i++ {
		<-done
	}
}
