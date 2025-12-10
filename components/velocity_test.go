package components

import (
	"math"
	"testing"
)

func TestVelocity_Mask(t *testing.T) {
	v := NewVelocity()
	if v.Mask() != MaskVelocity {
		t.Errorf("Velocity.Mask() = %v, want %v", v.Mask(), MaskVelocity)
	}
}

func TestVelocity_NewVelocity(t *testing.T) {
	v := NewVelocity()
	if v.X != 0 || v.Y != 0 {
		t.Errorf("NewVelocity() = (%v, %v), want (0, 0)", v.X, v.Y)
	}
}

func TestVelocity_WithX(t *testing.T) {
	v := NewVelocity().WithX(50)
	if v.X != 50 {
		t.Errorf("Velocity.WithX(50).X = %v, want 50", v.X)
	}
}

func TestVelocity_WithY(t *testing.T) {
	v := NewVelocity().WithY(75)
	if v.Y != 75 {
		t.Errorf("Velocity.WithY(75).Y = %v, want 75", v.Y)
	}
}

func TestVelocity_With(t *testing.T) {
	v := NewVelocity().With(30, 40)
	if v.X != 30 || v.Y != 40 {
		t.Errorf("Velocity.With(30, 40) = (%v, %v), want (30, 40)", v.X, v.Y)
	}
}

func TestVelocity_Magnitude(t *testing.T) {
	tests := []struct {
		name string
		x, y float32
		want float32
	}{
		{"zero", 0, 0, 0},
		{"unit x", 1, 0, 1},
		{"unit y", 0, 1, 1},
		{"negative x", -1, 0, 1},
		{"3-4-5 triangle", 3, 4, 5},
		{"negative 3-4", -3, -4, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewVelocity().With(tt.x, tt.y)
			got := v.Magnitude()
			if math.Abs(float64(got-tt.want)) > 0.001 {
				t.Errorf("Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVelocity_Chaining(t *testing.T) {
	v := NewVelocity().WithX(10).WithY(20)
	if v.X != 10 || v.Y != 20 {
		t.Errorf("Velocity chaining failed: got (%v, %v), want (10, 20)", v.X, v.Y)
	}
}
