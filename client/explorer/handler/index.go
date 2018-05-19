package handler

import (
	"html/template"
	"net/http"

	"github.com/bsandusky/protocache/client/cache"
)

// Index serves template with cache details
func Index(w http.ResponseWriter, r *http.Request) {
	res, _ := cache.GetAll()
	tmpl := template.Must(template.ParseFiles("client/explorer/public/index.html"))
	tmpl.Execute(w, res)
}
