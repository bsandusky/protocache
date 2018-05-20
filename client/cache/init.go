// Package cache provides wrappers for gRPC client methods in order to offer a simple API to client applications
package cache

import (
	"fmt"

	"github.com/bsandusky/protocache/pb"
	"google.golang.org/grpc"
)

var (
	client pb.CacheClient
)

// InitClient establishes connection with gRPC server and returns error if not initialized
func InitClient(host, port string) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client = pb.NewCacheClient(conn)
	fmt.Printf("Cache client initialized\n")
	return nil
}
