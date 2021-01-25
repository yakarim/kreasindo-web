package config

import (
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger"
)

// HTML ...
func (c *Config) HTML(ctx *atreugo.RequestCtx, code int, page string, data H) error {

	t, vars := c.template(ctx, code, page)
	return t.Execute(ctx.RequestCtx, vars, data)
}

func (c *Config) template(ctx *atreugo.RequestCtx, code int, page string) (*jet.Template, jet.VarMap) {
	t, err := views.GetTemplate(page + ".html")
	if err != nil {
		logger.Fatal(err)
	}
	c.globalFunc(views)
	vars := make(jet.VarMap)

	ctx.SetStatusCode(code)
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	return t, vars
}

func (c *Config) globalFunc(view *jet.Set) {
	tm, _ := c.TimeIn(time.Now(), "Local")
	view.AddGlobal("time", tm.Format("01/02/2006"))
}
