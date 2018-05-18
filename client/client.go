package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		switch args[0] {
		case "get":
			res, err := client.Get(context.Background(), &pb.GetRequest{Key: args[1]})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
		case "set":
			res, err := client.Set(context.Background(), &pb.SetRequest{Key: args[1], Value: []byte(strings.Join(args[2:], " "))})
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
			res, err := client.FlushKey(context.Background(), &pb.FlushKeyRequest{Key: args[1]})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Command not recognized")
		}
	}

	if scanner.Err() != nil {
		// handle error.
	}
}
