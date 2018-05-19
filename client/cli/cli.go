package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bsandusky/protocache/client/cache"
)

// Start runs cli client
func Start() {
	var (
		res map[string]string
		err error
	)

	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")

		switch args[0] {
		case "get":
			res, err = cache.Get(args[1])
		case "getall":
			res, err = cache.GetAll()
		case "set":
			res, err = cache.Set(args[1], strings.Join(args[2:], " "))
		case "flushall":
			res, err = cache.FlushAll()
		case "flushkey":
			res, err = cache.FlushKey(args[1])
		case "exit":
			os.Exit(0)
		default:
			err = errors.New("Command not recognized")
		}
		handleOutput(res, err)
	}

	if scanner.Err() != nil {
		// handle error.
	}
}

func handleOutput(res map[string]string, err error) {

	if err != nil {
		fmt.Println(err)
		fmt.Printf("> ")
		return
	}

	fmt.Printf("< %s\n", res)
	fmt.Printf("> ")
}
