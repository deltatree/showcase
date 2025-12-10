package components

// Color represents the RGBA color of an entity with gradient support.
type Color struct {
	R, G, B, A uint8
	// For gradient interpolation
	StartR, StartG, StartB, StartA uint8
	EndR, EndG, EndB, EndA         uint8
}

// Mask returns the component mask for Color.
func (c *Color) Mask() uint64 { return MaskColor }

// NewColor creates a new Color component with default white color.
func NewColor() *Color {
	return &Color{
		R: 255, G: 255, B: 255, A: 255,
		StartR: 255, StartG: 255, StartB: 255, StartA: 255,
		EndR: 255, EndG: 255, EndB: 255, EndA: 0,
	}
}

// WithRGBA sets the RGBA values and initializes the gradient start.
func (c *Color) WithRGBA(r, g, b, a uint8) *Color {
	c.R, c.G, c.B, c.A = r, g, b, a
	c.StartR, c.StartG, c.StartB, c.StartA = r, g, b, a
	return c
}

// WithGradient sets the start and end colors for gradient interpolation.
func (c *Color) WithGradient(sr, sg, sb, sa, er, eg, eb, ea uint8) *Color {
	c.StartR, c.StartG, c.StartB, c.StartA = sr, sg, sb, sa
	c.EndR, c.EndG, c.EndB, c.EndA = er, eg, eb, ea
	c.R, c.G, c.B, c.A = sr, sg, sb, sa
	return c
}
