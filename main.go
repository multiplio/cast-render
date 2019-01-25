package main

import (
	"flag"
	"log"

	env "github.com/Netflix/go-env"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"

	"github.com/tekwrks/renderer/render"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "fonts/NotoSans-Regular.ttf", "filename of the ttf font")
	address  = flag.String("address", ":3000", "http service address")
	name     = flag.String("alias", "renderer", "program name")
)

type environment struct {
	IPFSAddress string `env:"IPFS_ADDRESS"`
}

func main() {
	// get command line options
	flag.Parse()

	// setup logger
	log.SetFlags(0)
	log.SetPrefix(*name + ":")

	// get environment
	var environment environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	// init renderer
	renderer, err := render.InitRenderer(fontfile, dpi)
	if err != nil {
		log.Fatal("Could not init render with fontfile", *fontfile, " : ", err)
		return
	}

	// renderer context struct
	renderContext := &renderContext{
		renderer: &renderer,
		shell:    shell.NewShell(environment.IPFSAddress),
	}

	// routes
	router := routing.New()
	setRoutes(router, renderContext)

	// start server
	log.Println("Serving at ", *address)
	log.Fatal(fasthttp.ListenAndServe(*address, router.HandleRequest))
}
