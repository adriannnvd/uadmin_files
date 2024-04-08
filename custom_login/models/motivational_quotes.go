package models

import (
	"github.com/uadmin/uadmin"
)

// Motivational Quotes Model !
type MotivationalQuotes struct {
	uadmin.Model
	Name    string `uadmin:"display_name:Type"`
	Author string 
	Quote   string
	Picture string `uadmin:"image"`
	Like    string
	Comment string
	Repost  string
}
