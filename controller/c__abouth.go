package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
	"github.com/yakarim/kreasindo-web/model"
)

// Abouth ...
type Abouth struct {
	config.Config
	model.Model
}

// AbouthDepan view.
func (c *Abouth) AbouthDepan(ctx *atreugo.RequestCtx) error {
	data, _ := c.Abouth.Query()
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/abouth", config.H{
		"title":    "Tentang Kami",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
	})
}

// View abouth.
func (c *Abouth) View(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "abouth/view", config.H{
		"title":    "Tentang Kami admin",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

// JSON abouth.
func (c *Abouth) JSON(ctx *atreugo.RequestCtx) error {
	data, _ := c.Abouth.Query()
	return ctx.JSONResponse(config.H{"abouth": data}, 200)
}

// Create abouth.
func (c *Abouth) Create(ctx *atreugo.RequestCtx) error {
	var d database.Abouth

	title := ctx.FormValue("title")
	desc := ctx.FormValue("desc")

	if string(title) == "" || string(desc) == "" {
		ctx.JSONResponse(config.H{"msg": "isi form"}, 404)
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		database.DB.First(&d)
		d.Title = string(title)
		d.Desc = string(desc)
		return c.Abouth.Create(d)
	}

	nameimg, body, types, size := c.Image(file, 500, 0)
	database.DB.First(&d)
	d.Title = string(title)
	d.Desc = string(desc)
	d.Image = database.Images{Name: nameimg, Body: body, Type: types, Size: size}

	return c.Abouth.Create(d)
}
