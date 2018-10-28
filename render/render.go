package render

// Copyright 2010 The Freetype-Go Authors. All rights reserved.
// Use of this source code is governed by the GNU General Public License
// version 2 (or any later version), which can be found in the LICENSE file.

import (
	"golang.org/x/image/font"
	"image"
	"image/draw"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Renderer struct {
	font *truetype.Font
	dpi  float64
	fg   *image.Uniform
	bg   *image.Uniform
}

func InitRenderer(fontfile *string, dpi *float64) (Renderer, error) {
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		return Renderer{}, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return Renderer{}, err
	}

	r := Renderer{}
	r.font = f
	r.dpi = *dpi
	r.fg, r.bg = image.Black, image.White

	return r, nil
}

func (r *Renderer) Render(text []string, size *float64, spacing *float64) (*image.RGBA, error) {
	//Twitter requirements:
	//Aspect ratio of 2:1 with minimum dimensions of 300x157 or maximum of 4096x4096 pixels.
	//Images must be less than 5MB in size. JPG, PNG, WEBP and GIF formats are supported.
	scale := 3
	imageWidth := 300 * scale
	imageHeight := 157 * scale

	rgba := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(rgba, rgba.Bounds(), r.bg, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(r.dpi)
	c.SetFont(r.font)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(r.fg)
	c.SetHinting(font.HintingNone)

	// Draw the text.
	pt := freetype.Pt(0, (imageHeight-textHeight(text, c, size, spacing))/2)
	for _, s := range text {
		w, err := textWidth(s, r.font, size)
		if err != nil {
			return image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight)), err
		}
		pt.X = c.PointToFixed(float64(imageWidth-w) / 2)

		_, err = c.DrawString(s, pt)
		if err != nil {
			return image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight)), err
		}
		pt.Y += c.PointToFixed(*size * *spacing)
	}

	return rgba, nil
}
