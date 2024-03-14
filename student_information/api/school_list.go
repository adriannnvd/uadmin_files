package api

import (
	"net/http"

	"github.com/adriannnvd/student_information/models"
	"github.com/uadmin/uadmin"
)

// TodoListAPIHandler !
func SchoolListAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all records in the database
	school := []models.School{}
	results := []map[string]interface{}{}
	uadmin.All(&school)

	// Accesses and fetches data from another model
	for i := range school {
		uadmin.Preload(&school[i])
		// Assigns the string of interface in each Todo fields
		results = append(results, map[string]interface{}{
			"ID":      school[i].ID,
			"Name":    school[i].Name,
			"Logo":    school[i].Logo,
			"Address": school[i].Address,
			"Website": school[i].Website,
		})
	}

	// Return todo JSON object
	uadmin.ReturnJSON(w, r, results)
}
