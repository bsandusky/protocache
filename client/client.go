package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bsandusky/protocache/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// Establish client connection to server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewCacheClient(conn)

	flag.Parse()
	switch flag.Arg(0) {
	case "get":
		res, err := client.Get(context.Background(), &pb.GetRequest{Key: flag.Arg(1)})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	case "set":
		res, err := client.Set(context.Background(), &pb.SetRequest{Key: flag.Arg(1), Value: []byte(flag.Arg(2))})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	case "flushall":
		res, err := client.FlushAll(context.Background(), &pb.Empty{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	case "flushkey":
		res, err := client.FlushKey(context.Background(), &pb.FlushKeyRequest{Key: flag.Arg(1)})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	default:
	}
}
