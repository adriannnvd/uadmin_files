package api

import (
	"net/http"

	"github.com/adriannnvd/student_information/models"
	"github.com/uadmin/uadmin"
)

// TodoListAPIHandler !
func CourseListAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all records in the database
	course := []models.Course{}
	results := []map[string]interface{}{}
	uadmin.All(&course)

	// Accesses and fetches data from another model
	for i := range course {
		uadmin.Preload(&course[i])
		// Assigns the string of interface in each Todo fields
		results = append(results, map[string]interface{}{
			"ID":          course[i].ID,
			"Name":        course[i].Name,
			"Course Code": course[i].CourseCode,
		})
	}

	// Return todo JSON object
	uadmin.ReturnJSON(w, r, results)
}
