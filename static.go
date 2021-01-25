package main

import (
	"os"
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
		CacheDuration:   15 * time.Hour,
		Compress:        true,
	}
	ctx.StaticCustom("/"+name, rootFS)
}

func portj() (atreugo.Config, string) {
	port := os.Getenv("PORT")
	var config atreugo.Config
	if port == "8080" {
		config = atreugo.Config{
			Addr:              ":" + port,
			Name:              "Kreasindo Pratama",
			ReduceMemoryUsage: true,
			Compress:          true,
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
	return config, port
}
