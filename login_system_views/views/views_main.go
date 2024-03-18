package views

import (
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/logim_system")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	LoginHandler(w, r)
}