package controller

import (
	"log"

	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/model"
)

// Gallery ...
type Gallery struct {
	config.Config
	model.Model
}

// GalleryDepan ...
func (c *Gallery) GalleryDepan(ctx *atreugo.RequestCtx) error {
	data, err := c.Gallery.Query()
	if err != nil {
		log.Fatal("gallery err", err)
	}
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/gallery", config.H{
		"title":    "Gallery",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
	})
}

// View person.
func (c *Gallery) View(ctx *atreugo.RequestCtx) error {

	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "gallery/view", config.H{
		"title":    "Gallery",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

// JSON gallery.
func (c *Gallery) JSON(ctx *atreugo.RequestCtx) error {

	data, err := c.Gallery.Query()
	if err != nil {
		log.Fatal("JSON gallery error", err)
	}

	return ctx.JSONResponse(config.H{"contacts": data}, 200)

}

// Create person.
