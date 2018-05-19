package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

// FlushAll removes all keys and values from cache
func FlushAll() (map[string]string, error) {
	data := make(map[string]string)
	res, err := client.FlushAll(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}
	data["status"] = res.Result
	return data, nil
}

// FlushKey removes one key, value pair from cache
func FlushKey(key string) (map[string]string, error) {
	data := make(map[string]string)
	res, err := client.FlushKey(context.Background(), &pb.FlushKeyRequest{Key: []byte(key)})
	if err != nil {
		return nil, err
	}
	data["status"] = res.Result
	return data, nil
}
