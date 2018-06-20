package main

import (
	"github.com/WhoMeNope/notatweet/textrender"

	"bufio"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "fonts/ClearSans-Regular.ttf", "filename of the ttf font")
	size     = flag.Float64("size", 32, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
)

var text = []string{
	"’Twas brillig, and the slithy toves",
	"Did gyre and gimble in the wabe;",
	"All mimsy were the borogoves,",
	"And the mome raths outgrabe.",
}

func main() {
	flag.Parse()

	renderer, err := textrender.InitRenderer(fontfile, dpi, size, spacing)
	if err != nil {
		log.Println("Could not init textrenderer with fontfile", fontfile, " : ", err)
		return
	}

	// render
	rgba, err := renderer.Render(text)
	if err != nil {
		log.Println("Could not render text", text, " : ", err)
		return
	}

	// Save that RGBA image to disk.
	err = writeToFile("out.png", rgba)
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("Wrote out.png OK.")
	}
}

func writeToFile(filename string, image *image.RGBA) error {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()

	b := bufio.NewWriter(outFile)
	err = png.Encode(b, image)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return err
}
