package main

import (
	"github.com/WhoMeNope/notatweet/textrender"

	"flag"
	"github.com/valyala/fasthttp"
	"log"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "fonts/NotoSans-Regular.ttf", "filename of the ttf font")
	size     = flag.Float64("size", 42, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	address  = flag.String("address", ":3000", "http service address")
)

var text = []string{
	"Arguing that you don't care about the",
	"right to privacy because you have",
	"nothing to hide is no different from",
	"saying you don't care about free speech",
	"because you have nothing to say.",
	"- Edward Snowden",
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	renderer, err := textrender.InitRenderer(fontfile, dpi, size, spacing)
	if err != nil {
		log.Fatal("Could not init textrenderer with fontfile", *fontfile, " : ", err)
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
