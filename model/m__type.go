package model

import (
	"github.com/yakarim/kreasindo-web/database"
)

// Model ...
type Model struct {
	User       User
	Contact    Contact
	Abouth     Abouth
	Specialist Specialist
	Gallery    Gallery
}

var db = database.DB
