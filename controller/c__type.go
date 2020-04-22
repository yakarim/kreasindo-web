package controller

import (
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
)

// Controller type.
type Controller struct {
	config.Config
	User    User
	Contact Contact
}

var db = database.DB
