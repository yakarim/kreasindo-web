package model

import (
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
)

// Model ...
type Model struct {
	config.Config
	database.User
}

var db = database.DB
