package cache

import (
	"context"
	"strings"

	"github.com/bsandusky/protocache/pb"
)

// Set assigns a key, value pair in the cache
func Set(key, value string) (map[string]string, error) {
	res, err := client.Set(context.Background(), &pb.SetRequest{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	data["status"] = res.Result
	return data, nil
}

// HSet assigns a key, multi-value object in the cache
func HSet(key, value string) (map[string]string, error) {
	val := strings.Replace(value, " ", ":", -1)
	res, err := client.Set(context.Background(), &pb.SetRequest{
		Key:   key,
		Value: val,
	})
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	data["status"] = res.Result
	return data, nil
}
