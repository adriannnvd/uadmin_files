package api

// import (
// 	"net/http"

// 	"github.com/adriannnvd/student_information/models"
// 	"github.com/uadmin/uadmin"
// )

// // TodoListAPIHandler !
// func StudentInfoListAPIHandler(w http.ResponseWriter, r *http.Request) {
// 	// Fetch all records in the database
// 	studentInfo := []models.StudentInfo{}
// 	results := []map[string]interface{}{}
// 	uadmin.All(&studentInfo)

// 	// Accesses and fetches data from another model
// 	// Loop to fetch the record of todo
// 	for i := range studentInfo {
// 		// Accesses and fetches the record of the linking models in Todo
// 		uadmin.Preload(&studentInfo[i])

// 		// Assigns the string of interface in each Todo fields
// 		results = append(results, map[string]interface{}{
// 			"ID":             studentInfo[i].ID,
// 			"Name":           studentInfo[i].Name,
// 			"Student Number": studentInfo[i].StudentNo,
// 			"Address":        studentInfo[i].Address,
// 			"Contact Number": studentInfo[i].Contact,
// 			"Email":          studentInfo[i].Email,
// 			"Birthdate":      studentInfo[i].Birthdate,
// 			"Age":            studentInfo[i].Age,
// 			"Gender":         studentInfo[i].Gender,

// 			"Guardian": studentInfo[i].Guardian.Name,

// 			"School": studentInfo[i].School.Name,

// 			"Course":     studentInfo[i].Course,
// 			"Year Level": studentInfo[i].YearLevel,
// 		})
// 	}
// 	// Return todo JSON object
// 	uadmin.ReturnJSON(w, r, results)
// }
