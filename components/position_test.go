package components

import (
	"testing"
)

func TestPosition_Mask(t *testing.T) {
	p := NewPosition()
	if p.Mask() != MaskPosition {
		t.Errorf("Position.Mask() = %v, want %v", p.Mask(), MaskPosition)
	}
}

func TestPosition_NewPosition(t *testing.T) {
	p := NewPosition()
	if p.X != 0 || p.Y != 0 {
		t.Errorf("NewPosition() = (%v, %v), want (0, 0)", p.X, p.Y)
	}
}

func TestPosition_WithX(t *testing.T) {
	p := NewPosition().WithX(100)
	if p.X != 100 {
		t.Errorf("Position.WithX(100).X = %v, want 100", p.X)
	}
}

func TestPosition_WithY(t *testing.T) {
	p := NewPosition().WithY(200)
	if p.Y != 200 {
		t.Errorf("Position.WithY(200).Y = %v, want 200", p.Y)
	}
}

func TestPosition_With(t *testing.T) {
	p := NewPosition().With(100, 200)
	if p.X != 100 || p.Y != 200 {
		t.Errorf("Position.With(100, 200) = (%v, %v), want (100, 200)", p.X, p.Y)
	}
}

func TestPosition_Chaining(t *testing.T) {
	p := NewPosition().WithX(10).WithY(20)
	if p.X != 10 || p.Y != 20 {
		t.Errorf("Position chaining failed: got (%v, %v), want (10, 20)", p.X, p.Y)
	}
}
