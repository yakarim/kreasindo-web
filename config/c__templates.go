package config

import (
	"os"
	"time"

	"github.com/CloudyKit/jet/v3"
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger"
)

// HTML ...
func (c *Config) HTML(ctx *atreugo.RequestCtx, code int, page string, data H) error {
	p := os.Getenv("PORT")
	if p == "8080" {
		views.SetDevelopmentMode(true)
		t, vars := c.template(ctx, code, page)
		return t.Execute(ctx.RequestCtx, vars, data)
	}

	t, vars := c.template(ctx, code, page)
	return t.Execute(ctx.RequestCtx, vars, data)
}

func (c *Config) template(ctx *atreugo.RequestCtx, code int, page string) (*jet.Template, jet.VarMap) {
	views.Delims("[%", "%]")
	t, err := views.GetTemplate(page + ".html")
	if err != nil {
		logger.Fatal(err)
	}
	go c.globalFunc(views)
	vars := make(jet.VarMap)

	ctx.SetStatusCode(code)
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	return t, vars
}

func (c *Config) globalFunc(view *jet.Set) {
	tm, _ := c.TimeIn(time.Now(), "Local")
	view.AddGlobal("time", tm.Format("01/02/2006"))
}
