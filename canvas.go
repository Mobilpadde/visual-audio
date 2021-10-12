package visual

import (
	"image"
	"math"

	"github.com/fogleman/gg"
)

type Canvas struct {
	dc          *gg.Context
	samples     []int16
	spacing     int
	sampleWidth int
	height      int
}

func Blank(samples []int16, spacing, sampleWidth, height int) *Canvas {
	l := len(samples)
	w := float64(sampleWidth * l * spacing)
	h := float64(height)

	dc := gg.NewContext(int(w), height)

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

func (c *Canvas) Branding(path string, alpha float64, repeat bool) (*Canvas, error) {
	im, err := gg.LoadImage(path)
	if err != nil {
		return c, err
	}

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

func (c *Canvas) Waves(r, g, b int, padding float64) *Canvas {
	h := float64(c.height)
	ws := float64(c.sampleWidth)
	s := float64(c.spacing)

	min, max := minMax(c.samples)

	c.dc.SetRGB255(r, g, b)
	for i, sample := range c.samples {
		tmpHeight := mapNumber(
			float64(sample),
			min,
			max,
			padding,
			h-padding,
		)

		c.dc.DrawRoundedRectangle(float64(i)*ws*s+s, float64(c.height)/2-tmpHeight/2, ws, tmpHeight, ws/2)
		c.dc.Fill()
	}

	return c
}

func (c *Canvas) Save(out string) error {
	return c.dc.SavePNG(out)
}

func (c *Canvas) Image() image.Image {
	return c.dc.Image()
}

func minMax(vals []int16) (float64, float64) {
	min, max := int16(math.MaxInt16), int16(math.MinInt16)
	for _, x := range vals {
		if min > x {
			min = x
		}

		if max < x {
			max = x
		}
	}

	return float64(min), float64(max)
}

func mapNumber(x, inMin, inMax, outMin, outMax float64) float64 {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
