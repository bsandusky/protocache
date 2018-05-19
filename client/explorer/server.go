package explorer

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bsandusky/protocache/client/explorer/handler"
)

// Start runs web explorer client
func Start() {

	r := http.NewServeMux()
	r.HandleFunc("/", handler.Index)

	srv := &http.Server{
		Addr:         ":5000",
		Handler:      r,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	fmt.Printf("Cache explorer running at %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
