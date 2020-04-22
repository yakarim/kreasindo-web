package config

import (
	"github.com/savsgio/atreugo/v10"
	"github.com/valyala/fasthttp"
)

// Auth login
func (c *Config) Auth(ctx *atreugo.RequestCtx) (*UserCredential, bool, error) {
	var signIn bool
	jwtCookie := ctx.Request.Header.Cookie("__jwt")
	_, user, err := validateToken(string(jwtCookie))
	if err != nil {
		signIn = false
	} else {
		signIn = true
	}
	return user, signIn, err
}

// AuthMiddleware  ..
func (c *Config) AuthMiddleware(ctx *atreugo.RequestCtx) error {
	if string(ctx.Path()) == "/login" {
		return ctx.Next()
	}
	jwtCookie := ctx.Request.Header.Cookie("__jwt")

	if len(jwtCookie) == 0 {
		return ctx.RedirectResponse("/login", fasthttp.StatusForbidden)
	}
	token, _, err := validateToken(string(jwtCookie))
	if err != nil {
		return err
	}
	if !token.Valid {
		return ctx.RedirectResponse("/login", fasthttp.StatusForbidden)
	}
	return ctx.Next()
}

// SecurityTime ...
func (c *Config) SecurityTime(ctx *atreugo.RequestCtx) error {
	ctx.Response.Header.Set("X-Frame-Options", "SAMEORIGIN")
	ctx.Response.Header.Set("X-XSS-Protection", "1; mode=block")
	ctx.Response.Header.Set("Strict-Transport-Security", "max-age= max-age=30672000")
	ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")
	return ctx.Next()
}
