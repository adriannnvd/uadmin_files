package main

import (
	"net/http"

	"github.com/uadmin/uadmin"

	"github.com/adriannnvd/uadmin_files/custom_login/models"
	"github.com/adriannnvd/uadmin_files/custom_login/views"
)

func main() {
	// Assign RootURL value as "/admin/" and Site Name as "Login System"
	// NOTE: This code works only if database does not exist yet.
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Quotes of the Day"

	uadmin.Register(
		models.RandomQuotes{},
		models.InspirationalQuotes{},
		models.MotivationalQuotes{},
		models.LoveQuotes{},
		models.PerceveranceQuotes{},
		models.EncouragementQuotes{},
	)

	http.HandleFunc("/", uadmin.Handler(views.MainHandler))

	// Run the server
	uadmin.StartServer()
}
