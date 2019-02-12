package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"strings"

	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/qiangxue/fasthttp-routing"

	"github.com/tekwrks/renderer/render"
)

type renderContext struct {
	shell    *ipfs.Shell
	renderer *render.Renderer
}

func (r *renderContext) handleTwitter(c *routing.Context) error {
	hash := c.Param("hash")
	log.Println("Got hash :", hash)

	imagePath := environment.RootURL + "/post/" + hash + "/image"

	page := `
	<!doctype HTML>
		<head>
			<meta name="twitter:card" content="summary_large_image" />
			<meta name="og:title" content="Hello World" />
			<meta name="og:image" content="` + imagePath + `">
		</head>
		<body>
			<img alt="" src="` + imagePath + `" />
			Cast by tekwrks
		</body>
	</html>
	`
	c.Response.Header.SetContentType("text/html; charset=utf-8")
	fmt.Fprintf(c, page)
	return nil
}

type post struct {
	Description string  `json:"description"`
	Content     string  `json:"content"`
	FontSize    float64 `json:"fontSize"`
	Spacing     float64 `json:"spacing"`
}

func (r *renderContext) handleImage(c *routing.Context) error {
	hash := c.Param("hash")
	log.Println("Got hash :", hash)

	// get post from ipfs
	block, err := r.shell.BlockGet("/ipfs/" + hash)
	if err != nil {
		log.Println("error:", err)
		return routing.NewHTTPError(400, "No post found.")
	}
	var post post
	err = json.Unmarshal(block, &post)
	if err != nil {
		log.Println("error:", err)
		return routing.NewHTTPError(400, "Not a post.")
	}

	// render
	lines := strings.Split(post.Content, "\n")
	rgba, err := r.renderer.Render(lines, &post.FontSize, &post.Spacing)
	if err != nil {
		log.Println("Could not render : ", hash, " : ", err)
		return routing.NewHTTPError(400, "Not a post.")
	}

	// write
	err = png.Encode(c, rgba)
	if err != nil {
		log.Println(err)
		return routing.NewHTTPError(400, "Not a post.")
	}

	return nil
}
