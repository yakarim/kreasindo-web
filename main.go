package main

import (
	"runtime"

	"github.com/savsgio/atreugo/v11"
	fastprefork "github.com/valyala/fasthttp/prefork"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config, port := portj()
	ctx := atreugo.New(config)
	routers(ctx)
	static(ctx)

	if port == "8080" {
		if err := ctx.ListenAndServe(); err != nil {
			panic(err)
		}
	}
	preforkServer := &fastprefork.Prefork{
		RecoverThreshold: runtime.GOMAXPROCS(0) / 2,
		ServeFunc:        ctx.Serve,
	}

	if err := preforkServer.ListenAndServe(":" + port); err != nil {
		panic(err)
	}

}
