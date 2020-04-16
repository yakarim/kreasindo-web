package main

import (
	"time"

	"github.com/savsgio/atreugo/v10"
)

func static(ctx *atreugo.Atreugo) {

	ctx.Static("/css", "./static/css")
	ctx.Static("/js", "./static/js")
	rootFS := &atreugo.StaticFS{
		Root:            "./static/images",
		AcceptByteRange: true,
		CacheDuration:   15 * time.Minute,
	}
	ctx.StaticCustom("/images", rootFS)

}
