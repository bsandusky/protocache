package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

func FlushAll() (string, error) {
	res, err := client.FlushAll(context.Background(), &pb.Empty{})
	if err != nil {
		return "", err
	}
	return res.Status, nil
}

func FlushKey(key string) (string, error) {
	res, err := client.FlushKey(context.Background(), &pb.FlushKeyRequest{Key: key})
	if err != nil {
		return "", err
	}
	return res.Status, nil
}
