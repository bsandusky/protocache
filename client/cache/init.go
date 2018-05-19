package cache

import "github.com/bsandusky/protocache/pb"

var client pb.CacheClient

// InitCache sets client cache for cache package
func InitCache(c pb.CacheClient) {
	client = c
}
