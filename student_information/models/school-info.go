package models

import (
	"github.com/uadmin/uadmin"
)

// SchoolInfo Model !
type School struct {
	uadmin.Model
	Name        string `uadmin:"required;search"`
	SchoolCode  string `uadmin:"required"`
	Logo        string `uadmin:"required;image"`
	Address     string `uadmin:"required"`
	WebsiteLink string `uadmin:"required;list_exclude"`
	Website     string `uadmin:"link;website"`
}

func (s School) Validate() (errMsg map[string]string) {

	errMsg = map[string]string{}

	school := School{}
	if uadmin.Count(&school, "name = ? AND id <> ?", s.Name, s.ID) != 0 {
		errMsg["Name"] = "This school is already in the system"
	}
	return
}

func (sch *School) Save() {
	link := sch.WebsiteLink
	sch.Website = link
	uadmin.Save(sch)
}
