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
	ctx.GET("/specialist", specialist)
	ctx.GET("/contact", c.ContactDepan)
	ctx.GET("/login", c.Login)
	ctx.POST("/login__jwt", c.LoginJwt)

	ctxUser := ctx.NewGroupPath("/user-admin")
	ctxUser.UseBefore(c.AuthMiddleware)
	ctxUser.GET("", c.User.View)
	ctxUser.GET("/json", c.User.JSON)
	ctxUser.POST("", c.User.Create)
	ctxUser.PUT("", c.User.Update)
	ctxUser.DELETE("/del:key", c.User.Delete)

	ctxContact := ctx.NewGroupPath("/contact-admin")
	ctxContact.UseBefore(c.AuthMiddleware)
	ctxContact.GET("", c.Contact.View)

}

// Index ...
func Index(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)

	return c.HTML(ctx, 200, "pages/home", config.H{
		"title":    "Home",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

func specialist(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	return c.HTML(ctx, 200, "pages/specialist", config.H{
		"title":    "Specialist",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}
