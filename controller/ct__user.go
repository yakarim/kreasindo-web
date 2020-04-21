package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
	"github.com/yakarim/kreasindo-web/model"
)

// User controller.
type User struct {
	model.Model
	config.Config
}

// View user ...
func (c *User) View(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	conf := config.H{
		"title":    "User Pages",
		"username": string(u.Username),
		"signIn":   signIn,
	}
	return c.HTML(ctx, 200, "users/view", conf)
}

// JSON user.
func (c *User) JSON(ctx *atreugo.RequestCtx) error {
	data, _ := c.User.Query()
	return ctx.JSONResponse(config.H{"users": data}, 200)
}

// Create user.
func (c *User) Create(ctx *atreugo.RequestCtx) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if string(username) == "" || string(email) == "" || string(password) == "" {
		ctx.JSONResponse(config.H{"msg": "form harus disi"}, 404)
	} else {
		var user database.User
		user.Username = string(username)
		user.Email = string(email)
		user.Password = string(password)
		err := c.User.Create(user)
		if err != nil {
			ctx.JSONResponse(config.H{"msg": "Email Found"}, 404)
		} else {
			ctx.JSONResponse(config.H{"msg": "sukses"}, 200)
		}
	}
	return nil
}

// Update user ...
func (c *User) Update(ctx *atreugo.RequestCtx) error {
	var user database.User
	ID := ctx.FormValue("ID")
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user.ID = string(ID)
	user.Username = string(username)
	user.Email = string(email)
	user.Password = string(password)

	err := c.User.Update(user)
	if err != nil {
		ctx.JSONResponse(config.H{"msg": err.Error()}, 404)
	} else {
		ctx.JSONResponse(config.H{"msg": "sukses"}, 200)
	}
	return nil
}

// Delete user.
func (c *User) Delete(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	err := c.User.Delete(key.(string))
	if err != nil {
		ctx.JSONResponse(config.H{"msg": err.Error()}, 404)
	}
	ctx.Redirect("/user", 200)

	return nil
}
