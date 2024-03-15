package main

import (
	// "net/http"

	"github.com/adriannnvd/student_information/models"

	// "github.com/adriannnvd/student_information/api"

	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.StudentInfo{},
		models.School{},
		models.Guardian{},
	)

	uadmin.RegisterInlines(models.School{}, map[string]string{
		"StudentInfo": "SchoolID",
	})

	uadmin.RegisterInlines(models.Guardian{}, map[string]string{
		"StudentInfo": "GuardianID",
	})

	// http.HandleFunc("/api/", uadmin.Handler(api.Handler))

	uadmin.StartServer()

}
