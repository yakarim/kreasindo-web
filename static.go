package main

import (
	"time"

	"github.com/savsgio/atreugo/v11"
)

func static(ctx *atreugo.Atreugo) {

	costumStatic(ctx, "css")
	costumStatic(ctx, "js")
	costumStatic(ctx, "images")

}

func costumStatic(ctx *atreugo.Atreugo, name string) {
	rootFS := &atreugo.StaticFS{
		Root:            "./static/" + name,
		AcceptByteRange: true,
		CacheDuration:   15 * time.Minute,
	}
	ctx.StaticCustom("/"+name, rootFS)
}
