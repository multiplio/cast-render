package main

import (
	"github.com/tekwrks/renderer/render"

	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "fonts/NotoSans-Regular.ttf", "filename of the ttf font")
	address  = flag.String("address", ":3000", "http service address")
	name     = flag.String("alias", "renderer", "program name")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix(*name + ":")

	renderer, err := render.InitRenderer(fontfile, dpi)
	if err != nil {
		log.Fatal("Could not init render with fontfile", *fontfile, " : ", err)
		return
	}

	// renderer context struct
	h := &renderHandler{
		renderer: &renderer,
	}

	// request handler
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/render":
			h.handleRender(ctx)
		default:
			ctx.Error("404 : path not found", fasthttp.StatusNotFound)
		}
	}
	log.Println("Serving at ", *address)
	log.Fatal(fasthttp.ListenAndServe(*address, requestHandler))
}
