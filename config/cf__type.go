package config

import "github.com/CloudyKit/jet/v3"

// Config kreasindo.
type Config struct{}

var (
	views = jet.NewHTMLSet("./template")
)
