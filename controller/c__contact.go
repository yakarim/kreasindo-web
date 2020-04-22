package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
)

// Contact ...
type Contact struct {
	config.Config
}

// ContactDepan view.
func (c *Controller) ContactDepan(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/contact", config.H{
		"title":    "Contact",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

// View contact.
func (c *Contact) View(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "contact/view", config.H{
		"title":    "Contact",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}
