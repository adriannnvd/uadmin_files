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

	randomCount := uadmin.Count(&randomQuotes, "id > 1")
	// uadmin.Trail(uadmin.DEBUG, "randomCount: %v\n", randomCount)
	randomMinID := 2
	randomMaxID := randomCount
	randomID := randomizeNumber(randomMinID, randomMaxID)
	uadmin.Get(&c.RandomQuote, "id = ?", randomID)

	//For Inspirational Quotes
	inspirationalQuotes := []models.InspirationalQuotes{}

	inspirationalCount := uadmin.Count(&inspirationalQuotes, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "inspirationalCount: %v\n", inspirationalCount)
	inspirationalMinID := 1
	inspirationalMaxID := inspirationalCount
	inspirationalID := randomizeNumber(inspirationalMinID, inspirationalMaxID)
	uadmin.Get(&c.InspirationalQuote, "id = ?", inspirationalID)

	// For Motivational Quotes
	motivationalQuotes := []models.MotivationalQuotes{}

	motivationalCount := uadmin.Count(&motivationalQuotes, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "motivationalCount: %v\n", motivationalCount)
	motivationalMinID := 1
	motivationalMaxID := motivationalCount
	motivationalID := randomizeNumber(motivationalMinID, motivationalMaxID)
	uadmin.Get(&c.MotivationalQuote, "id = ?", motivationalID)

	// For Love Quotes
	loveQuotes := []models.LoveQuotes{}

	loveCount := uadmin.Count(&loveQuotes, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "loveCount: %v\n", loveCount)
	loveMinID := 1
	loveMaxID := loveCount
	loveID := randomizeNumber(loveMinID, loveMaxID)
	uadmin.Get(&c.LoveQuote, "id = ?", loveID)

	// For Perceverance Quotes
	perseveranceQuotes := []models.PerceveranceQuotes{}

	perseveranceCount := uadmin.Count(&perseveranceQuotes, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "perseveranceCount: %v\n", perseveranceCount)
	perseveranceMinID := 1
	perseveranceMaxID := perseveranceCount
	perseveranceID := randomizeNumber(perseveranceMinID, perseveranceMaxID)
	uadmin.Get(&c.PerceveranceQuote, "id = ?", perseveranceID)

	// For Encouragement Quotes
	encouragementQuotes := []models.EncouragementQuotes{}

	encouragementCount := uadmin.Count(&encouragementQuotes, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "encouragementCount: %v\n", encouragementCount)
	encouragementMinID := 1
	encouragementMaxID := encouragementCount
	encouragementID := randomizeNumber(encouragementMinID, encouragementMaxID)
	uadmin.Get(&c.EncouragementQuote, "id = ?", encouragementID)

	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}

// randomizeNumber generates a random number within the specified range [min, max].
func randomizeNumber(min, max int) int {

	// Generate a random number within the specified range.

	randmized := rand.Intn(max-min+1) + min

	return randmized
}
