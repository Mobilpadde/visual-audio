package visual

import (
	"image"

	"github.com/fogleman/gg"
)

// Branding takes an image (im) to then provide a branded canvas
func (c *Canvas) Branding(im image.Image, alpha float64, repeat bool) (*Canvas, error) {
	if repeat {
		size := im.Bounds().Max.X
		_, max := minMax([]int16{int16(size), int16(c.dc.Width())})
		for i := 0; i < int(max); i += size {
			c.dc.DrawImageAnchored(im, i, c.dc.Height()/2, 0, 0.5)
		}
	} else {
		c.dc.DrawImageAnchored(im, c.dc.Width()/2, c.dc.Height()/2, 0.5, 0.5)
	}

	c.dc.DrawRectangle(0, 0, float64(c.dc.Width()), float64(c.dc.Height()))
	c.dc.SetRGBA(1, 1, 1, alpha)
	c.dc.Fill()

	return c, nil
}

// BrandingPath takes a path to an image to then provide a branded canvas
func (c *Canvas) BrandingPath(path string, alpha float64, repeat bool) (*Canvas, error) {
	im, err := gg.LoadImage(path)
	if err != nil {
		return c, err
	}

	return c.Branding(im, alpha, repeat)
}
