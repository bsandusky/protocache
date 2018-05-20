package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bsandusky/protocache/client/cache"
)

// Start runs cli client
func Start(done chan bool) {
	var (
		key string
		val string
		res map[string]string
		err error
	)
	re := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)

	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		args := re.FindAllString(scanner.Text(), -1)

		if len(args) >= 2 {
			key = trimQuotes(args[1])
			val = trimQuotes(strings.Join(args[2:], " "))
		}

		switch args[0] {
		case "get":
			res, err = cache.Get(key)
		case "getall":
			res, err = cache.GetAll()
		case "set":
			res, err = cache.Set(key, val)
		case "hset":
			res, err = cache.HSet(key, val)
		case "flushall":
			res, err = cache.FlushAll()
		case "flushkey":
			res, err = cache.FlushKey(key)
		case "exit":
			done <- true
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

func trimQuotes(s string) string {
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, "'", "", -1)
	return s
}

func handleOutput(res map[string]string, err error) {

	if err != nil {
		fmt.Printf("%v\n", err)
		fmt.Printf("> ")
		return
	}

	fmt.Printf("< %s\n", res)
	fmt.Printf("> ")
}
