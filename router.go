package main

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/controller"
)

var c controller.Controller

func routers(ctx *atreugo.Atreugo) {
	ctx.GET("/", Index)
	ctx.GET("/user", c.UserTemplates)

}

// Index ...
func Index(ctx *atreugo.RequestCtx) error {
	return c.HTML(ctx, 200, "pages/index", config.H{
		"title": "Halaman Depan",
	})
}
