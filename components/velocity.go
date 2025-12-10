package components

import "math"

// Velocity represents the movement speed and direction of an entity.
// It is measured in pixels per second and applied by the physics system.
//
// The X component represents horizontal movement (positive = right),
// and Y represents vertical movement (positive = down).
type Velocity struct {
	// X is the horizontal velocity in pixels per second.
	X float32
	// Y is the vertical velocity in pixels per second.
	Y float32
}

// Mask returns the component mask for Velocity.
func (v *Velocity) Mask() uint64 { return MaskVelocity }

// NewVelocity creates a new Velocity component.
func NewVelocity() *Velocity { return &Velocity{} }

// WithX sets the X velocity and returns the velocity for chaining.
func (v *Velocity) WithX(x float32) *Velocity { v.X = x; return v }

// WithY sets the Y velocity and returns the velocity for chaining.
func (v *Velocity) WithY(y float32) *Velocity { v.Y = y; return v }

// With sets both X and Y velocity and returns the velocity for chaining.
func (v *Velocity) With(x, y float32) *Velocity { v.X = x; v.Y = y; return v }

// Magnitude calculates the length of the velocity vector.
func (v *Velocity) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}
