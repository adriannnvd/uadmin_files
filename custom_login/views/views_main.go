package views

import (
	"net/http"
	"strings"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/login_system/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

    // LoginHandler verifies login data and creating sessions for users.
    LoginHandler(w, r)
}
