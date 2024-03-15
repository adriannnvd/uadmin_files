package models

import (
	"github.com/uadmin/uadmin"
)

// SchoolInfo Model !
type School struct {
	uadmin.Model
	Name string `uadmin:"required;search"`
	Logo       string `uadmin:"required;image"`
	Address    string `uadmin:"required"`
	Website    string `uadmin:"required"`
}
