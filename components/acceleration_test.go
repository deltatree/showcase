package components

import (
	"testing"
)

func TestAcceleration_Mask(t *testing.T) {
	a := NewAcceleration()
	if a.Mask() != MaskAcceleration {
		t.Errorf("Acceleration.Mask() = %v, want %v", a.Mask(), MaskAcceleration)
	}
}

func TestAcceleration_NewAcceleration(t *testing.T) {
	a := NewAcceleration()
	if a.X != 0 || a.Y != 0 {
		t.Errorf("NewAcceleration() = (%v, %v), want (0, 0)", a.X, a.Y)
	}
}

func TestAcceleration_Reset(t *testing.T) {
	a := NewAcceleration()
	a.X = 100
	a.Y = 200
	a.Reset()
	if a.X != 0 || a.Y != 0 {
		t.Errorf("After Reset(): (%v, %v), want (0, 0)", a.X, a.Y)
	}
}

func TestAcceleration_Add(t *testing.T) {
	a := NewAcceleration()
	a.Add(10, 20)
	if a.X != 10 || a.Y != 20 {
		t.Errorf("After Add(10, 20): (%v, %v), want (10, 20)", a.X, a.Y)
	}

	a.Add(5, 10)
	if a.X != 15 || a.Y != 30 {
		t.Errorf("After Add(5, 10): (%v, %v), want (15, 30)", a.X, a.Y)
	}
}

func TestAcceleration_WithX(t *testing.T) {
	a := NewAcceleration().WithX(50)
	if a.X != 50 {
		t.Errorf("Acceleration.WithX(50).X = %v, want 50", a.X)
	}
}

func TestAcceleration_WithY(t *testing.T) {
	a := NewAcceleration().WithY(75)
	if a.Y != 75 {
		t.Errorf("Acceleration.WithY(75).Y = %v, want 75", a.Y)
	}
}

func TestAcceleration_Chaining(t *testing.T) {
	a := NewAcceleration().WithX(10).WithY(20)
	if a.X != 10 || a.Y != 20 {
		t.Errorf("Acceleration chaining failed: got (%v, %v), want (10, 20)", a.X, a.Y)
	}
}
