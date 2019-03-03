package main

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
)

func setRoutes(router *routing.Router, renderContext *renderContext) {

	router.Get("/ready", func(context *routing.Context) error {
		fmt.Fprintf(context, "ok")
		return nil
	})

	router.Get("/post/<hash>", func(context *routing.Context) error {
		return renderContext.handleTwitter(context)
	})

	router.Get("/post/<hash>/image", func(context *routing.Context) error {
		return renderContext.handleImage(context)
	})

}
