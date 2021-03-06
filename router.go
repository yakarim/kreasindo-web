package main

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/controller"
	"github.com/yakarim/kreasindo-web/database"
)

var c controller.Controller

func routers(ctx *atreugo.Atreugo) {

	ctx.UseBefore(c.SecurityTime)
	ctx.GET("/", Index)
	ctx.GET("/specialist", c.Specialist.SpecialistDepan)
	ctx.GET("/contact", c.Contact.ContactDepan)
	ctx.GET("/abouth", c.Abouth.AbouthDepan)
	ctx.GET("/gallery", c.Gallery.View)
	ctx.GET("/person-json", c.Gallery.JSON)

	ctx.GET("/login", c.Login)
	ctx.POST("/login__jwt", c.LoginJwt)
	ctx.GET("/images", Image)
	ctx.GET("/imagesv/:id", c.ShowImage)

	ctxUser := ctx.NewGroupPath("/user-admin")
	ctxUser.UseBefore(c.AuthMiddleware)
	ctxUser.GET("", c.User.View)
	ctxUser.GET("/json", c.User.JSON)
	ctxUser.POST("", c.User.Create)
	ctxUser.PUT("", c.User.Update)
	ctxUser.DELETE("/del{key}", c.User.Delete)

	ctxContact := ctx.NewGroupPath("/contact-admin")
	ctxContact.UseBefore(c.AuthMiddleware)
	ctxContact.GET("", c.Contact.View)
	ctxContact.GET("/json", c.Contact.JSON)
	ctxContact.POST("", c.Contact.Create)

	ctxAbouth := ctx.NewGroupPath("/abouth-admin")
	ctxAbouth.UseBefore(c.AuthMiddleware)
	ctxAbouth.GET("", c.Abouth.View)
	ctxAbouth.GET("/json", c.Abouth.JSON)
	ctxAbouth.POST("", c.Abouth.Create)

	ctxSpecialist := ctx.NewGroupPath("/specialist-admin")
	ctxSpecialist.UseBefore(c.AuthMiddleware)
	ctxSpecialist.GET("", c.Specialist.View)
	ctxSpecialist.GET("/json", c.User.JSON)
	ctxSpecialist.POST("", c.User.Create)
	ctxSpecialist.PUT("", c.User.Update)
	ctxSpecialist.DELETE("/del{key}", c.User.Delete)

	ctxGallery := ctx.NewGroupPath("/gallery-admin")
	ctxGallery.UseBefore(c.AuthMiddleware)
	ctxGallery.GET("", c.Gallery.View)
	ctxGallery.GET("/json", c.Gallery.JSON)
}

// Image ...
func Image(ctx *atreugo.RequestCtx) error {
	var img []database.Images
	database.DB.Find(&img)
	return ctx.JSONResponse(img, 200)
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
