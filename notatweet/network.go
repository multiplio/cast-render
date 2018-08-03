package main

import (
	"github.com/WhoMeNope/notatweet/textrender"

	"github.com/valyala/fasthttp"
	"image/png"
	"log"
	"os"
	"strings"
)

type renderHandler struct {
	renderer *textrender.Renderer
}

func (h *renderHandler) handleRender(ctx *fasthttp.RequestCtx) {
	var (
		text    []string
		size    float64
		spacing float64
		err     error
	)

	// get params
	args := ctx.QueryArgs()
	log.Println("Parsing query args : ", args)

	if args.Has("text") {
		rawtext := string(args.Peek("text"))
		text = strings.Split(rawtext, "\n")
	}
	size, err = args.GetUfloat("size")
	if err != nil {
		log.Println("Could not get params : ", err)
		return
	}
	spacing, err = args.GetUfloat("spacing")
	if err != nil {
		log.Println("Could not get params : ", err)
		return
	}

	// render
	rgba, err := h.renderer.Render(text, &size, &spacing)
	if err != nil {
		log.Println("Could not render : ", text, " : ", err)
		return
	}

	// write
	err = png.Encode(ctx, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
