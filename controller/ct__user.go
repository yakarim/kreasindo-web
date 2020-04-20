package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/model"
)

// User controller.
type User struct {
	model.Model
	config.Config
}

// View user ...
func (c *User) View(ctx *atreugo.RequestCtx) error {

	u, signIn, _ := c.Auth(ctx)
	data, _ := c.User.Query()

	conf := config.H{
		"title":    "User Pages",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
	}

	return c.HTML(ctx, 200, "users/view", conf)
}
