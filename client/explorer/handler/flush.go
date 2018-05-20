package handler

import (
	"net/http"

	"github.com/bsandusky/protocache/client/cache"
)

// Flushkey clears one key-value pair from cache
func Flushkey(w http.ResponseWriter, r *http.Request) {
	cache.FlushKey(r.URL.Query().Get("key"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

// Flushall clears all key-value pairs from cache
func Flushall(w http.ResponseWriter, r *http.Request) {
	cache.FlushAll()
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
