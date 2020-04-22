package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
)

// Contact view.
func (c *Controller) Contact(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/contact", config.H{
		"title":    "Contact",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}
