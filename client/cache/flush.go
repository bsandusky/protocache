package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

// FlushAll removes all keys and values from cache
func FlushAll() (*pb.Result, error) {

	res, err := client.FlushAll(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FlushKey removes one key, value pair from cache
func FlushKey(key string) (*pb.Result, error) {
	res, err := client.FlushKey(context.Background(), &pb.FlushKeyRequest{Key: key})
	if err != nil {
		return nil, err
	}
	return res, nil
}
