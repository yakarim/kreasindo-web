package config

import (
	"time"

	"github.com/CloudyKit/jet/v3"
	"github.com/savsgio/atreugo/v10"
	"github.com/savsgio/go-logger"
)

// HTML ...
func (c *Config) HTML(ctx *atreugo.RequestCtx, code int, page string, data H) error {

	views.SetDevelopmentMode(true)
	views.Delims("[%", "%]")
	t, err := views.GetTemplate(page + ".jet.html")
	if err != nil {
		logger.Fatal(err)
	}
	go c.globalFunc(views)
	vars := make(jet.VarMap)

	ctx.SetStatusCode(code)
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")

	return t.Execute(ctx.RequestCtx, vars, data)

}

func (c *Config) globalFunc(view *jet.Set) {
	tm, _ := c.TimeIn(time.Now(), "Local")
	view.AddGlobal("time", tm.Format("01/02/2006"))
}
