package config

import (
	"time"

	"github.com/CloudyKit/jet/v3"
	"github.com/fasthttp-contrib/render"
	"github.com/savsgio/atreugo/v10"
)

// HTML tempates.
func (c *Config) HTML(ctx *atreugo.RequestCtx, code int, page string, data H) error {

	r := render.New(&render.Config{
		Directory:  "templates",                // Specify what path to load the templates from.
		Gzip:       true,                       // enable it if you want render through gzip compression
		Layout:     "layout/base",              // Specify a layout template. Layouts can call {{ yield }} to render the current template or {{ partial "css" }} to render a partial from the current template.
		Extensions: []string{".html", ".tpml"}, // Specify extensions to load for templates.
		Delims:     render.Delims{"[%", "%]"},  // Sets delimiters to the specified strings.
		// Output XHTML content type instead of default "text/html".
		IsDevelopment:             true, // Render will now recompile the templates on every HTML response.
		UnEscapeHTML:              true, // Replace ensure '&<>' are output correctly (JSON only).
		StreamingJSON:             true, // Streams the JSON response via json.Encoder.
		RequirePartials:           true, // Return an error if a template is missing a partial used in a layout.
		DisableHTTPErrorRendering: true, // Disables automatic rendering of http.StatusInternalServerError when an error occurs.
	})
	return r.HTML(ctx.RequestCtx, code, page, data)
}

// HTMLJet ...
func (c *Config) HTMLJet(ctx *atreugo.RequestCtx, code int, page string, data H) error {
	views.SetDevelopmentMode(true)
	views.Delims("[%", "%]")
	t, err := views.GetTemplate(page + ".jet.html")
	if err != nil {
		// template could not be loaded
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
