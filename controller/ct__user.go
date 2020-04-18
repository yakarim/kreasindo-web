package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
)

// UserTemplates ...
func (c *Controller) UserTemplates(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	data, _ := c.User.Query()
	return c.HTML(ctx, 200, "users/view", config.H{
		"title":    "User Pages",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
	})
}
