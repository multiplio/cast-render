package main

import (
	"fmt"
	// "image/png"
	"log"
	// "os"
	// "strings"

	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/valyala/fasthttp"

	"github.com/tekwrks/renderer/render"
)

type renderHandler struct {
	shell    *ipfs.Shell
	renderer *render.Renderer
}

func (h *renderHandler) handleRender(ctx *fasthttp.RequestCtx) {
	var (
		// text    []string
		// size    float64
		// spacing float64
		err error
	)

	// get params
	args := ctx.QueryArgs()
	log.Println("Parsing query args : ", args)

	if args.Has("id") {
		id := string(args.Peek("id"))
		log.Println("id:", id)
		ctx.SetStatusCode(fasthttp.StatusOK)
	} else {
		log.Println("Could not get params : ", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "error : provide required arguments")
		return
	}

	fmt.Fprintf(ctx, "ok")
	return

	// // render
	// rgba, err := h.renderer.Render(text, &size, &spacing)
	// if err != nil {
	// 	log.Println("Could not render : ", text, " : ", err)
	// 	return
	// }

	// // write
	// err = png.Encode(ctx, rgba)
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }
}
