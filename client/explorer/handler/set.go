package handler

import (
	"net/http"

	"github.com/bsandusky/protocache/client/cache"
)

// Set parses form and sets value in cache
func Set(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Invalid method"))
	}

	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	cache.Set(r.Form.Get("key"), r.Form.Get("value"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
