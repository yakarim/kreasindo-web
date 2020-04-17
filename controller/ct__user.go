package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
)

// UserTemplates ...
func (c *Controller) UserTemplates(ctx *atreugo.RequestCtx) error {
	data, _ := c.User.Query()
	str := false
	return c.HTML(ctx, 200, "backend/users/view", config.H{
		"title": "User Pages",
		"str":   str,
		"data":  data,
	})
}
