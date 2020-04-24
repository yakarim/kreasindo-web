package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/model"
)

// Specialist controller.
type Specialist struct {
	model.Model
	config.Config
}

// SpecialistDepan ...
func (c *Specialist) SpecialistDepan(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/specialist", config.H{
		"title":    "Specialist",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

// View specialist ...
func (c *Specialist) View(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "specialist/view", config.H{
		"title":    "Specialist Admin",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}
