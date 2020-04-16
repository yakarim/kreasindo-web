package main

import "github.com/savsgio/atreugo/v10"

func router(ctx *atreugo.Atreugo) {
	ctx.GET("/", Index)

}

// Index ...
func Index(ctx *atreugo.RequestCtx) error {
	return ctx.JSONResponse("halo walid", 200)
}
