package render

import (
	"fmt"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
)

func textWidth(text string, font *truetype.Font, fontSize *float64) (int, error) {
	opts := truetype.Options{}
	opts.Size = *fontSize
	face := truetype.NewFace(font, &opts)

	width := fixed.Int26_6(0)
	for _, x := range text {
		awidth, ok := face.GlyphAdvance(rune(x))
		if ok != true {
			return 0, fmt.Errorf("Glyph not found in the font")
		}
		width += awidth
	}
	return width.Ceil(), nil
}

func textHeight(text []string, context *freetype.Context, size *float64, spacing *float64) int {
	height := fixed.Int26_6(len(text)-1) * context.PointToFixed(*size**spacing)
	height -= context.PointToFixed(*size * *spacing / 2)
	return height.Ceil()
}
