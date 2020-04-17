package controller

import (
	"net/url"

	"github.com/savsgio/atreugo/v10"
	"github.com/valyala/fasttemplate"
)

// UserTemplates ...
func (c *Controller) UserTemplates(ctx *atreugo.RequestCtx) error {
	template := "[host]"
	t := fasttemplate.New(template, "[", "]")
	s := t.ExecuteString(map[string]interface{}{
		"host":  "google.com",
		"query": url.QueryEscape("hello=world"),
		"bar":   "foobar",
	})
	return ctx.HTTPResponseBytes([]byte(s), 200)
}
