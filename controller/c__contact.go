package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
	"github.com/yakarim/kreasindo-web/model"
)

// Contact ...
type Contact struct {
	config.Config
	model.Model
}

// ContactDepan view.
func (c *Contact) ContactDepan(ctx *atreugo.RequestCtx) error {
	data, _ := c.Contact.Query()
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/contact", config.H{
		"title":    "Contact",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
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

// JSON contact.
func (c *Contact) JSON(ctx *atreugo.RequestCtx) error {
	data, _ := c.Contact.Query()
	return ctx.JSONResponse(config.H{"contacts": data}, 200)
}

// Create contact.
func (c *Contact) Create(ctx *atreugo.RequestCtx) error {
	var d database.Contact

	telpon := ctx.FormValue("telpon")
	alamat := ctx.FormValue("alamat")

	if string(telpon) == "" || string(alamat) == "" {
		ctx.JSONResponse(config.H{"msg": "isi form"}, 404)
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		database.DB.First(&d)
		d.Alamat = string(alamat)
		d.Telp = string(telpon)
		return c.Contact.Create(d)
	}

	nameimg, body, types, size := c.Image(file, 500, 0)
	database.DB.First(&d)
	d.Alamat = string(alamat)
	d.Telp = string(telpon)
	d.Image = database.Images{Name: nameimg, Body: body, Type: types, Size: size}

	return c.Contact.Create(d)
}
