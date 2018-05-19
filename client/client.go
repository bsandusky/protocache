package main

import (
	"log"

	"github.com/bsandusky/protocache/client/cache"
	"github.com/bsandusky/protocache/client/cli"
	"github.com/bsandusky/protocache/client/explorer"
	"github.com/bsandusky/protocache/pb"
	"google.golang.org/grpc"
)

func main() {
	// Establish client connection to server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewCacheClient(conn)
	cache.InitCache(client)

	done := make(chan bool)
	go cli.Start(done)
	go explorer.Start()

	// Wait for exit command to close
	<-done

}
