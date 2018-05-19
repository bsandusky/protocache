package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

// FlushAll removes all keys and values from cache
func FlushAll() (map[string]string, error) {

	res, err := client.FlushAll(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	data["status"] = res.Result
	return data, nil
}

// FlushKey removes one key, value pair from cache
func FlushKey(key string) (map[string]string, error) {
	res, err := client.FlushKey(context.Background(), &pb.FlushKeyRequest{Key: key})
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	data["status"] = res.Result
	return data, nil
}
