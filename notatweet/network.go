package main

import (
	"github.com/WhoMeNope/notatweet/textrender"

	"github.com/valyala/fasthttp"
	"image/png"
	"log"
	"os"
)

type renderHandler struct {
	renderer *textrender.Renderer
}

func (h *renderHandler) handleRender(ctx *fasthttp.RequestCtx) {
	// render
	rgba, err := h.renderer.Render(text)
	if err != nil {
		log.Println("Could not render text", text, " : ", err)
		return
	}

	// write
	err = png.Encode(ctx, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
