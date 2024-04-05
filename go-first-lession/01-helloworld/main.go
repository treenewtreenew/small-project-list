package main

// In the whole go project module, there should be only one main package
// the main func is the bigining of all go programe
// Go is using UTF-8 charset
// We can compile go file by using go build ...., and run compiling result
// also we can execute go directly using go run ....

import (
	"fmt" // this fmt is the import path

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main() {
	// The formal syntax of go specification uses semicolons ';' as end markers.
	// Most semocolons are optional and are often omitted
	// Even if semicolons are added manually, gofmt will remove them automatically when formatting
	// the code according to conventions
	fmt.Println("你好，世界") // this fmt is package name

	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}
