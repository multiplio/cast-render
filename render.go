package main

import (
	"github.com/multiplio/cast-render/render"

	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"strings"

	"github.com/cbroglie/mustache"
	cid "github.com/ipfs/go-cid"
	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/qiangxue/fasthttp-routing"
)

type renderContext struct {
	shell        *ipfs.Shell
	renderer     *render.Renderer
	postTemplate *string
}

func (r *renderContext) handleTwitter(c *routing.Context) error {
	hash := c.Param("hash")
	log.Println("Got hash :", hash)

	imagePath := environment.RootURL + "/post/" + hash + "/image"

	page, err := mustache.Render(*r.postTemplate, map[string]string{
		"imagePath": imagePath,
		"title":     "cast by multipl",
	})
	if err != nil {
		log.Println("error:", err)
		return routing.NewHTTPError(400, "Could not render post.")
	}

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
	// get and validate hash
	hash := c.Param("hash")
	if hash == "" {
		log.Println("No hash in url")
		return routing.NewHTTPError(400, "Provide a post hash.")
	}

	// Create a cid from a marshaled string
	cid, err := cid.Decode(hash)
	if err != nil {
		log.Println("Hash is not valid cid")
		return routing.NewHTTPError(400, "Provide a valid post hash.")
	}
	log.Println("Got cid :", cid)

	// get post from ipfs
	block, err := r.shell.BlockGet("/ipfs/" + cid.String())
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
		log.Println("Could not render : ", cid, " : ", err)
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
