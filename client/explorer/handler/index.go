package handler

import (
	"net/http"

	"github.com/bsandusky/protocache/client/cache"
)

func Index(w http.ResponseWriter, r *http.Request) {
	res, _ := cache.GetAll()
	w.Write([]byte(res.String()))
}
