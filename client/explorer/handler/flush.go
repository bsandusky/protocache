package handler

import (
	"net/http"

	"github.com/bsandusky/protocache/client/cache"
)

// Flushkey clears one key from cache
func Flushkey(w http.ResponseWriter, r *http.Request) {
	cache.FlushKey(r.URL.Query().Get("key"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Flushall clears entire cache
func Flushall(w http.ResponseWriter, r *http.Request) {
	cache.FlushAll()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
