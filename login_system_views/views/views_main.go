package views

import (
    "net/http"
    "strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
    // r.URL.Path creates a new path called "/login_system/"
    r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login_system")
    r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

    // Authentication : This session is preloaded with a user.
    session := uadmin.IsAuthenticated(r)
    if session == nil {
        // LoginHandler verifies login data and creating sessions for users.
        LoginHandler(w, r)
        return
    }

    // HomeHandler handles the home page.
    HomeHandler(w, r, session)
    return
}