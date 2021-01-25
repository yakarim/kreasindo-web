package config

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/yakarim/kreasindo-web/database"
)

// Config kreasindo.
type Config struct {
}

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates"),
	jet.InDevelopmentMode(), // remove in production
	jet.WithDelims("[%", "%]"),
)
var (
	db = database.DB
)
