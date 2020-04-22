package model

import (
	"github.com/yakarim/kreasindo-web/database"
)

// Model ...
type Model struct {
	User    User
	Contact Contact
}

var db = database.DB
