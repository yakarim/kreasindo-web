package main

import (
	"os"
	"runtime"

	"github.com/savsgio/atreugo/v10"
	"github.com/savsgio/go-logger"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	port := os.Getenv("PORT")
	var config atreugo.Config
	if port == "8080" {
		config = atreugo.Config{
			Addr:     "0.0.0.0:8080",
			Name:     "Kreasindo Pratama",
			Compress: true,
			LogLevel: logger.DEBUG,
		}
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
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
	go routers(ctx)
	go static(ctx)

	if err := ctx.ListenAndServe(); err != nil {
		panic(err)
	}
}
