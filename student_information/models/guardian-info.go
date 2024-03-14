package models

import (
	"github.com/uadmin/uadmin"
)

// Guardian Model !
type Guardian struct {
	uadmin.Model
	Name          string `uadmin:"required;help:Firstname Middlename Lastname"`
	Relationship1 string `uadmin:"required;display_name:Relationship;help:Mother ,Father etc.."`
	Contact1      string `uadmin:"required;default_value: ;display_name:Contact;pattern:^[ ,0-9]*$;pattern_msg:Your input must be a number."`
	Email1        string `uadmin:"display_name:Email"`
	Address1      string `uadmin:"required;display_name:Address"`
	Occupation1   string `uadmin:"display_name:Occupation"`

	OtherGuardian string `uadmin:"display_name:Other Guardian;help:Firstname Middlename Lastname"`
	Relationship2 string `uadmin:"display_name:Relationship;help:Mother ,Father etc.."`
	Contact2      string `uadmin:"required;default_value: ;display_name:Contact;pattern:^[ ,0-9]*$;pattern_msg:Your input must be a number."`
	Email2        string `uadmin:"display_name:Email"`
	Address2      string `uadmin:"display_name:Address"`
	Occupation2   string `uadmin:"display_name:Occupation"`
}
