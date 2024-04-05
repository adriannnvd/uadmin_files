package views

import (
	"net/http"

	"math/rand"
	"time"

	"github.com/adriannnvd/uadmin_files/custom_login/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {

	type Context struct {
		User  string
		Quote models.Quotes
		Time  string
		Date  string
	}

	// Call the custom struct and assign the full name in the User field under the context object.
	c := Context{}
	// uadmin.Trail(uadmin.DEBUG, "session USER: "+session.User.FirstName+" "+session.User.LastName)
	c.User = session.User.FirstName + " " + session.User.LastName

	currentTime := time.Now()

	// Format time in AM/PM format
	c.Time = currentTime.Format("03:04 PM")

	// Format date in MM/DD/YYYY format
	c.Date = currentTime.Format("01/02/2006")

	// Print time in AM/PM format

	// If you want to assign these values to variables

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 50
	randomNumber := rand.Intn(9) + 1

	uadmin.Get(&c.Quote, "id = ?", randomNumber)

	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}
