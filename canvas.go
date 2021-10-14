package visual

import (
	"image"

	"github.com/fogleman/gg"
)

// Canvas is to contain data for our funcs
type Canvas struct {
	dc          *gg.Context
	samples     []int16
	spacing     int
	sampleWidth int
	height      int
}

// Blank will return a new *Canvas
func Blank(samples []int16, spacing, sampleWidth, height int, square bool) *Canvas {
	l := len(samples)
	w := float64(sampleWidth * l * spacing)
	h := float64(height)

	dc := gg.NewContext(int(w), height)
	if square {
		dc = gg.NewContext(height, height)
	}

	dc.DrawRectangle(0, 0, w, h)
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	return &Canvas{
		dc:          dc,
		samples:     samples,
		spacing:     spacing,
		sampleWidth: sampleWidth,
		height:      height,
	}
}

// Save will save the image to a specific path (out)
func (c *Canvas) Save(out string) error {
	return c.dc.SavePNG(out)
}

// Image will provide the image
func (c *Canvas) Image() image.Image {
	return c.dc.Image()
}
