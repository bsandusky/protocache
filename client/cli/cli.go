package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bsandusky/protocache/client/cache"
)

func Start() {
	var (
		res string
		err error
	)

	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")

		switch args[0] {
		case "get":
			res, err = cache.Get(args[1])
		case "set":
			res, err = cache.Set(args[1], strings.Join(args[2:], " "))
		case "flushall":
			res, err = cache.FlushAll()
		case "flushkey":
			res, err = cache.FlushKey(args[1])
		case "exit":
			os.Exit(0)
		default:
			res = "Command not recognized"
		}
		handleOutput(res, err)
	}

	if scanner.Err() != nil {
		// handle error.
	}
}

func handleOutput(res string, err error) {

	if err != nil {
		fmt.Println(err)
		fmt.Printf("> ")
		return
	}

	if res == "" {
		fmt.Println("Key not found")
		fmt.Printf("> ")
		return
	}

	fmt.Printf("< %s\n", res)
	fmt.Printf("> ")
}
