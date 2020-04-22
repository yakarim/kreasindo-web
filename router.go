package main

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/controller"
)

var c controller.Controller

func routers(ctx *atreugo.Atreugo) {
	ctx.UseBefore(c.SecurityTime)
	ctx.GET("/", Index)
	ctx.GET("/news", home)
	ctx.GET("/contact", c.Contact)
	ctx.GET("/login", c.Login)
	ctx.POST("/login__jwt", c.LoginJwt)

	ctxUser := ctx.NewGroupPath("/user")
	ctxUser.UseBefore(c.AuthMiddleware)
	ctxUser.GET("", c.User.View)
	ctxUser.GET("/json", c.User.JSON)
	ctxUser.POST("", c.User.Create)
	ctxUser.PUT("", c.User.Update)
	ctxUser.DELETE("/del:key", c.User.Delete)

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
