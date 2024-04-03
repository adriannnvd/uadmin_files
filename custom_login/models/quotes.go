package models

import (
	"github.com/uadmin/uadmin"
)

// Quotes Model !
type Quotes struct {
	uadmin.Model
	Name    string
	Quote   string `uadmin:"html"`
	Picture string `uadmin:"image"`
}
