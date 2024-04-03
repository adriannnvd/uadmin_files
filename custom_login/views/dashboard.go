package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {

	type Context struct {
        User        string
    }

    // Call the custom struct and assign the full name in the User field under the context object.
    c := Context{}
    c.User = session.User.FirstName + " " + session.User.LastName


    // Render the home filepath and pass the context data object to the HTML file.
    uadmin.RenderHTML(w, r, "templates/home.html", c)
    return

}
