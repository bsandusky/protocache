package cache

import "github.com/bsandusky/protocache/pb"

var client pb.CacheClient

func InitCache(c pb.CacheClient) {
	client = c
}
