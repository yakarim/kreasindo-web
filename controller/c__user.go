package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
)

// UserTemplates ...
func (c *Controller) UserTemplates(ctx *atreugo.RequestCtx) error {
	str := false
	return c.HTML(ctx, 200, "pages/index", config.H{
		"title": "User Pages",
		"str":   str,
	})
}