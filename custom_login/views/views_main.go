package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/quotes_of_the_day/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	path := strings.TrimSuffix(r.URL.Path, "/")

	// uadmin.Trail(uadmin.DEBUG, "Path: %s", path)
	session := uadmin.IsAuthenticated(r)

	switch path {
	case "login":
		LoginHandler(w, r)
	case "dashboard":
		// uadmin.Trail(uadmin.DEBUG, "session : "+session.User.FirstName+" "+session.User.LastName)
		DashboardHandler(w, r, session)
	case "logout":
		LogoutHandler(w, r, session)
	default:
		w.Write([]byte("Not Found"))
	}

	// // Authentication : This session is preloaded with a user.

	// if r.URL.Path == "/dashboard" {
	// 	// DasboardHandler handles the dashboard page.
	// 	DashboardHandler(w, r, session)
	// 	return
	// } else
	// if r.URL.Path == "/logout" {
	// 	/* If the request URL Path is /logout after the /login_system/, it will proceed to this part.
	// 	   e.g. localhost:8080/login_system/logout */

	// 	// LogoutHandler handles the logout process for the user.
	// 	LogoutHandler(w, r, session)
	// 	return

	// }
}
