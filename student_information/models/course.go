package models

import (
	"github.com/uadmin/uadmin"
)

// Course Info Model !
type Course struct {
	uadmin.Model
	Name       string `uadmin:"required;search;help: ex. BS Computer Engineering"`
	CourseCode string `uadmin:"required;help: BS CpE"`
}

func (s Course) Validate() (errMsg map[string]string) {

	errMsg = map[string]string{}

	course := Course{}
	if uadmin.Count(&course, "name = ? AND id <> ?", s.Name, s.ID) != 0 {
		errMsg["Name"] = "This course is already in the system"
	}
	return
}
