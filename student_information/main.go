package main

import (
	"github.com/adriannnvd/student_information/models"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.StudentInfo{},
		models.School{},
	)

	uadmin.RegisterInlines(models.School{}, map[string]string{
		"StudentInfo": "SchoolID",
	})

	uadmin.StartServer()

}
