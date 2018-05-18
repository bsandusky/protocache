package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

func Get(key string) (string, error) {
	res, err := client.Get(context.Background(), &pb.GetRequest{Key: key})
	if err != nil {
		return "", err
	}
	return string(res.Value), nil
}
