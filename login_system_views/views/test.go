package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func TestHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	type Context struct {
		User string
	}
	c := Context{}
	c.User = session.User.FirstName + " " + session.User.LastName

	uadmin.RenderHTML(w, r, "templates/test.html", c)
	return

}
