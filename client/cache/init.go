// Package cache provides wrappers for gRPC client methods in order to offer a simple API to client applications
package cache

import (
	"fmt"

	"github.com/bsandusky/protocache/pb"
)

var (
	client pb.CacheClient
)

// InitCache sets client cache for cache package
func InitCache(cacheClient pb.CacheClient) {
	client = cacheClient
	fmt.Printf("Cache initialized\n")
}
