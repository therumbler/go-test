package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprintf(ctx, "👋 Hi there!  RequestURI is %q", ctx.RequestURI())
	case "/health":
		fmt.Fprintf(ctx, "🏥 Healthy!")
	case "/liveness":
		fmt.Fprintf(ctx, "🎸 Live!")
	default:
		ctx.Error("Not found", fasthttp.StatusNotFound)
	}
}

func main() {
	log.Print("listening on http://127.0.0.1:8080")
	fasthttp.ListenAndServe(":8080", fastHTTPHandler)
}
