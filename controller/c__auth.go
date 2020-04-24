package controller

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/valyala/fasthttp"
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
)

// Login pages.
func (c *Controller) Login(ctx *atreugo.RequestCtx) error {
	u, signIn, _ := c.Auth(ctx)
	deleteSession(ctx)
	return c.HTML(ctx, 200, "login", config.H{
		"title":    "Login",
		"username": string(u.Username),
		"signIn":   signIn,
	})
}

// LoginJwt ...
func (c *Controller) LoginJwt(ctx *atreugo.RequestCtx) error {

	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if len(email) == 0 || len(password) == 0 {
		ctx.RedirectResponse("/login", ctx.Response.StatusCode())
	}
	u, b := c.validateUser(string(email), string(password))
	if b == true {
		jwtCookie := ctx.Request.Header.Cookie("__jwt")

		if len(jwtCookie) == 0 {
			tokenString, expireAt := c.GenerateToken([]byte(u.Username), []byte(u.Password))
			// Set cookie for domain
			cookie := fasthttp.AcquireCookie()
			defer fasthttp.ReleaseCookie(cookie)
			cookie.SetKey("__jwt")
			cookie.SetValue(tokenString)
			cookie.SetExpire(expireAt)
			ctx.Response.Header.SetCookie(cookie)
		}
		ctx.RedirectResponse("/", ctx.Response.StatusCode())
	} else {
		ctx.RedirectResponse("/login", ctx.Response.StatusCode())
	}
	return ctx.Err()
}

func (c *Controller) validateUser(email, password string) (database.User, bool) {
	var user database.User
	if !db.Where("email = ?", email).First(&user).RecordNotFound() {
		pwd2 := c.GetPwd(password)
		pwdMatch := c.ComparePasswords(user.Password, pwd2)
		return user, pwdMatch
	}
	return user, false
}

func deleteSession(ctx *atreugo.RequestCtx) {
	ctx.Response.Header.DelAllCookies()
	ctx.Response.Header.DelClientCookie("__jwt")
}
