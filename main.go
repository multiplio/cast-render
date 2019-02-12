package main

import (
	"log"

	env "github.com/Netflix/go-env"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"

	"github.com/tekwrks/renderer/render"
)

type environmentDesc struct {
	Name        string `env:"NAME"`
	Address     string `env:"ADDRESS"`
	DPI         int    `env:"DPI"`
	FontFile    string `env:"FONTFILE"`
	IPFSAddress string `env:"IPFS_ADDRESS"`
	RootURL     string `env:"ROOT_URL"`
}

var environment environmentDesc

func main() {
	log.SetFlags(0)

	// get environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	log.SetPrefix(environment.Name + ":")

	// init renderer
	dpi := float64(environment.DPI)
	renderer, err := render.InitRenderer(&environment.FontFile, &dpi)
	if err != nil {
		log.Fatal("Could not init render with fontfile", environment.FontFile, " : ", err)
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
	log.Println("Serving at ", environment.Address)
	log.Fatal(fasthttp.ListenAndServe(environment.Address, router.HandleRequest))
}
