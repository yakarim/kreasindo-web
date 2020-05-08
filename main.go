package main

import (
	"os"
	"runtime"

	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger"
	fastprefork "github.com/valyala/fasthttp/prefork"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	port := os.Getenv("PORT")
	var config atreugo.Config
	if port == "8080" {
		port = "8080"
		config = atreugo.Config{
			Addr:              ":" + port,
			Name:              "Kreasindo Pratama",
			ReduceMemoryUsage: true,
			Compress:          true,
			LogLevel:          logger.DEBUG,
		}
	} else {
		config = atreugo.Config{
			Addr:              "0.0.0.0:" + port,
			Name:              "Kreasindo Pratama",
			ReduceMemoryUsage: true,
			Compress:          true,
			Concurrency:       3,
			GracefulShutdown:  true,
		}
	}

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
