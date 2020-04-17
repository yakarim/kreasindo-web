package controller

import (
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/model"
)

// Controller type.
type Controller struct {
	model.Model
	config.Config
}
