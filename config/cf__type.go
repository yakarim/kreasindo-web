package config

import (
	"github.com/CloudyKit/jet/v3"
	"github.com/yakarim/kreasindo-web/database"
)

// Config kreasindo.
type Config struct {
}

var (
	views = jet.NewHTMLSet("./templates")
	db    = database.DB
)
