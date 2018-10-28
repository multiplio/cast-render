package main

import (
	"bufio"
	"image"
	"image/png"
	"log"
	"os"
)

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
