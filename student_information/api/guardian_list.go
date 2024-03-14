package api

import (
	"net/http"

	"github.com/adriannnvd/student_information/models"
	"github.com/uadmin/uadmin"
)

// TodoListAPIHandler !
func GuardianListAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all records in the database
	guardian := []models.Guardian{}
	results := []map[string]interface{}{}
	uadmin.All(&guardian)

	// Accesses and fetches data from another model
	for i := range guardian {
		uadmin.Preload(&guardian[i])
		// Assigns the string of interface in each Todo fields
		results = append(results, map[string]interface{}{
			"ID":              guardian[i].ID,
			"Name":            guardian[i].Name,
			"Relationship":    guardian[i].Relationship1,
			"Contact Number":  guardian[i].Contact1,
			"Email":           guardian[i].Email1,
			"Address":         guardian[i].Address1,
			"Other Guardian":  guardian[i].Name,
			"Relationship1":   guardian[i].Relationship1,
			"Contact Number1": guardian[i].Contact1,
			"Email1":          guardian[i].Email1,
			"Address1":        guardian[i].Address1,
		})
	}

	// Return todo JSON object
	uadmin.ReturnJSON(w, r, results)
}
