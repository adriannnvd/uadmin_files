package models

import (
	"github.com/uadmin/uadmin"
)

// StudentInfo Model !
type StudentInfo struct {
	uadmin.Model
	Name            string `uadmin:"required;search;help:Firstname Middlename Lastname"`
	Address         string `uadmin:"required"`
	Contact         string `uadmin:"required;default_value: ;display_name:Contact Number;pattern:^[ ,0-9]*$;pattern_msg:Your input must be a number."`
	Age             int    `uadmin:"required;default_value: "`
	Gender          string `uadmin:"required"`
	Guardian        string `uadmin:"required;help:Firstname Middlename Lastname"`
	Guardian_Number string `uadmin:"required;default_value: ;display_name:Guardian Number;pattern:^[ ,0-9]*$;pattern_msg:Your input must be a number."`
	School          School `uadmin:"required;filter"`
	SchoolID        uint
	Course          string `uadmin:"required;filter;help:ex: BS Computer Engineering ..."`
	YearLevel       string `uadmin:"required;help:ex: BS 1st Year, 2nd Year ..."`
}

// // Save function !
// func (s *StudentInfo) Save() {
// 	// Save the model to DB
// 	uadmin.Save(s)
// 	// Some other business Logic ...
// }

// // Validate function !
// func (si *StudentInfo) Validate() (errMsg map[string]string) {
// 	// Initialize the error messages
// 	errMsg = map[string]string{}
// 	// Get any records from the database that maches the name of
// 	// this record and make sure the record is not the record we are
// 	// editing right now
// 	student_information := StudentInfo{}
// 	if uadmin.Count(&student_information, "name = ? AND id <> ?", si.Name, si.ID) != 0 {
// 		errMsg["Name"] = "This student is already in the system"
// 	}
// 	return
// }
