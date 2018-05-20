package handler

import (
	"html/template"
	"net/http"

	"github.com/bsandusky/protocache/client/cache"
)

// GetAll serves go template with cache data
func GetAll(w http.ResponseWriter, r *http.Request) {
	res, _ := cache.GetAll()
	// TODO: Handle error
	tmpl := template.Must(template.ParseFiles("client/explorer/public/index.html"))
	tmpl.Execute(w, res)
}
