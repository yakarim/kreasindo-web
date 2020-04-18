package main

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/controller"
)

var c controller.Controller

func routers(ctx *atreugo.Atreugo) {
	ctx.GET("/", Index)
	ctx.GET("/news", home)
	ctx.GET("/login", c.Login)
	ctx.POST("/login__jwt", c.LoginJwt)

	ctxUser := ctx.NewGroupPath("/user__view")
	ctxUser.UseBefore(c.AuthMiddleware)
	ctxUser.GET("", c.UserTemplates)

}

// Index ...
func Index(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)

	return c.HTML(ctx, 200, "home", config.H{
		"title":    "Halaman Depan",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

func home(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "home", config.H{
		"title":    "Halaman Depan",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}
