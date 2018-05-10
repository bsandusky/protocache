package main

import (
	"log"
	"net"

	"github.com/bsandusky/protocache/pb"
	"github.com/bsandusky/protocache/server/cache"
	"google.golang.org/grpc"
)

func main() {
	store := make(map[string][]byte)

	// Server implementation
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCacheServer(grpcServer, &cache.Server{Store: store})
	grpcServer.Serve(lis)
}
