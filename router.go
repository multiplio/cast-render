package main

import (
	"github.com/qiangxue/fasthttp-routing"
)

func setRoutes(router *routing.Router, renderContext *renderContext) {

	router.Get("/post/<hash>", func(context *routing.Context) error {
		return renderContext.handleTwitter(context)
	})

	router.Get("/post/<hash>/image", func(context *routing.Context) error {
		return renderContext.handleImage(context)
	})

}
