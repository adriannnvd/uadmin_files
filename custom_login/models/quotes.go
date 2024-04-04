package models

import (
	"github.com/uadmin/uadmin"
)

// Quotes Model !
type Quotes struct {
	uadmin.Model
	Name    string
	Quote   string
	Picture string `uadmin:"image"`
	Like    string
	Comment string
	Repost  string
}
