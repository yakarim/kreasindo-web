package model

import (
	"github.com/yakarim/kreasindo-web/database"
)

// Model ...
type Model struct {
	User User
}

var db = database.DB
