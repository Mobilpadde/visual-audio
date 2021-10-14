package visual

// Waves will generate a visual representation of the audio from samples
func (c *Canvas) Waves(r, g, b int, padding float64) *Canvas {
	h := float64(c.height)
	ws := float64(c.sampleWidth)
	s := float64(c.spacing)

	min, max := minMax(c.samples)

	c.dc.SetRGB255(r, g, b)
	for i, sample := range c.samples {
		hi := mapNumber(
			float64(sample),
			min,
			max,
			padding,
			h-padding,
		)

		c.dc.DrawRoundedRectangle(float64(i)*ws*s+s, float64(c.height)/2-hi/2, ws, hi, ws/2)
		c.dc.Fill()
	}

	return c
}
