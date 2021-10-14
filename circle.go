package visual

import "math"

// CircleWaves will generate a visual representation of the audio from samples in a circle
func (c *Canvas) CircleWaves(r, g, b int, padding float64) *Canvas {
	ws := float64(c.sampleWidth)
	s := float64(c.spacing)

	min, max := minMax(c.samples)
	rad := (360 / (float64(len(c.samples)) * ws * s)) / (math.Pi / 4)

	c.dc.SetRGB255(r, g, b)
	for i, sample := range c.samples {
		hi := mapNumber(
			float64(sample),
			min,
			max,
			10,
			float64(c.height)/2-padding,
		)

		c.dc.Push()
		c.dc.Translate(float64(c.dc.Width())/2, float64(c.height)/2)
		c.dc.Rotate(rad * float64(i))
		c.dc.DrawRoundedRectangle(0, padding, ws, hi, ws/2)
		c.dc.Fill()
		c.dc.Pop()

		if rad*float64(i) >= math.Pi*2 {
			break
		}
	}

	return c
}
