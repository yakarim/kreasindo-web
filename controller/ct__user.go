package controller

import (
	"fmt"

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
	data, _ := c.User.Query()

	conf := config.H{
		"title":    "User Pages",
		"username": string(u.Username),
		"signIn":   signIn,
		"data":     data,
	}

	return c.HTML(ctx, 200, "users/view", conf)
}

// Create user.
func (c *User) Create(ctx *atreugo.RequestCtx) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if string(username) == "" || string(email) == "" || string(password) == "" {
		fmt.Println("isi form")
		ctx.Redirect("./user", 200)
	} else {
		var user database.User
		user.Username = string(username)
		user.Email = string(email)
		user.Password = string(password)
		err := c.User.Create(user)
		if err != nil {
			ctx.Redirect("./user", 200)
			fmt.Println(err)
		}
		ctx.Redirect("./user", 200)
	}
	return nil
}

// Delete user.
func (c *User) Delete(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	err := c.User.Delete(key.(string))
	if err != nil {
		ctx.Redirect("/user", 200)
		fmt.Println(err)
	}
	ctx.Redirect("/user", 200)

	return nil
}
