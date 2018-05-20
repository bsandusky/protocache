package handler

import (
	"net/http"

	"github.com/bsandusky/protocache/client/cache"
)

// Set middleware routes values to basic or hash set storage
// cache.Set parses form data and adds or updates basic string-string key-value pairs in cache
// cache.HSet parses form data and adds or updates hashed string-multistring key-value pairs in cache
func Set(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Invalid method"))
	}

	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if r.Form.Get("hash") == "on" {
		cache.HSet(r.Form.Get("key"), r.Form.Get("value"))
	} else {
		cache.Set(r.Form.Get("key"), r.Form.Get("value"))
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
