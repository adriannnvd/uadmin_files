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
		User               string
		RandomQuote        models.RandomQuotes
		MotivationalQuote  models.MotivationalQuotes
		InspirationalQuote models.InspirationalQuotes
		PerceveranceQuote  models.PerceveranceQuotes
		LoveQuote          models.LoveQuotes
		EncouragementQuote models.EncouragementQuotes
		Time               string
		Date               string
	}

	// Call the custom struct and assign the full name in the User field under the context object.
	c := Context{}

	// Getting the user firstname and lastname
	c.User = session.User.FirstName + " " + session.User.LastName

	// getting the current time
	currentTime := time.Now()

	// Format time in AM/PM format
	c.Time = currentTime.Format("03:04 PM")

	// Format date in MM/DD/YYYY format
	c.Date = currentTime.Format("01/02/2006")

	// Count and Get the Quote in models
	// For Random Quotes
	randomQuotes := []models.RandomQuotes{}

	randomCount := uadmin.Count(&randomQuotes, "id > 0")
	uadmin.Trail(uadmin.DEBUG, "randomCount: %v\n", randomCount)
	uadmin.Get(&c.RandomQuote, "id = ?", randomizeNumber(randomCount))

	//For Inspirational Quotes
	inspirationalQuotes := []models.InspirationalQuotes{}

	inspirationalCount := uadmin.Count(&inspirationalQuotes, "id > 0")
	uadmin.Trail(uadmin.DEBUG, "motivationalCount: %v\n", inspirationalCount)
	uadmin.Get(&c.InspirationalQuote, "id = ?", randomizeNumber(inspirationalCount))

	// For Motivational Quotes
	motivationalQuotes := []models.MotivationalQuotes{}

	motivationalCount := uadmin.Count(&motivationalQuotes, "id > 0")
	uadmin.Trail(uadmin.DEBUG, "motivationalCount: %v\n", motivationalCount)
	uadmin.Get(&c.MotivationalQuote, "id = ?", randomizeNumber(motivationalCount))

	// For Love Quotes
	loveQuotes := []models.LoveQuotes{}

	loveCount := uadmin.Count(&loveQuotes, "id > 0")
	uadmin.Trail(uadmin.DEBUG, "motivationalCount: %v\n", loveCount)
	uadmin.Get(&c.LoveQuote, "id = ?", randomizeNumber(loveCount))

	// For Perceverance Quotes
	perceveranceQuotes := []models.LoveQuotes{}

	perceveranceCount := uadmin.Count(&perceveranceQuotes, "id > 0")
	uadmin.Trail(uadmin.DEBUG, "motivationalCount: %v\n", perceveranceCount)
	uadmin.Get(&c.PerceveranceQuote, "id = ?", randomizeNumber(perceveranceCount))

		// For Encouragement Quotes
		encouragementQuotes := []models.LoveQuotes{}

		encouragementCount := uadmin.Count(&encouragementQuotes, "id > 0")
		uadmin.Trail(uadmin.DEBUG, "motivationalCount: %v\n", encouragementCount)
		uadmin.Get(&c.EncouragementQuote, "id = ?", randomizeNumber(encouragementCount))

	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}

func randomizeNumber(number int) int {
	// Generate a random number between 0 and the given number
	randomized := rand.Intn(number) + 1

	return randomized
}
