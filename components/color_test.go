package components

import (
	"testing"
)

func TestColor_Mask(t *testing.T) {
	c := NewColor()
	if c.Mask() != MaskColor {
		t.Errorf("Color.Mask() = %v, want %v", c.Mask(), MaskColor)
	}
}

func TestColor_NewColor(t *testing.T) {
	c := NewColor()
	if c.R != 255 || c.G != 255 || c.B != 255 || c.A != 255 {
		t.Errorf("NewColor() = (%v, %v, %v, %v), want (255, 255, 255, 255)", c.R, c.G, c.B, c.A)
	}
}

func TestColor_WithRGBA(t *testing.T) {
	c := NewColor().WithRGBA(100, 150, 200, 128)
	if c.R != 100 || c.G != 150 || c.B != 200 || c.A != 128 {
		t.Errorf("WithRGBA() = (%v, %v, %v, %v), want (100, 150, 200, 128)", c.R, c.G, c.B, c.A)
	}
	if c.StartR != 100 || c.StartG != 150 || c.StartB != 200 || c.StartA != 128 {
		t.Errorf("WithRGBA() start colors = (%v, %v, %v, %v), want (100, 150, 200, 128)",
			c.StartR, c.StartG, c.StartB, c.StartA)
	}
}

func TestColor_WithGradient(t *testing.T) {
	c := NewColor().WithGradient(255, 0, 0, 255, 0, 0, 255, 0)

	if c.StartR != 255 || c.StartG != 0 || c.StartB != 0 || c.StartA != 255 {
		t.Errorf("WithGradient() start = (%v, %v, %v, %v), want (255, 0, 0, 255)",
			c.StartR, c.StartG, c.StartB, c.StartA)
	}

	if c.EndR != 0 || c.EndG != 0 || c.EndB != 255 || c.EndA != 0 {
		t.Errorf("WithGradient() end = (%v, %v, %v, %v), want (0, 0, 255, 0)",
			c.EndR, c.EndG, c.EndB, c.EndA)
	}

	if c.R != 255 || c.G != 0 || c.B != 0 || c.A != 255 {
		t.Errorf("WithGradient() current = (%v, %v, %v, %v), want (255, 0, 0, 255)",
			c.R, c.G, c.B, c.A)
	}
}
