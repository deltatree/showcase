package components

// Mass represents the gravitational mass of an entity.
// Mass affects how strongly attractors influence particle movement
// through the GravitySystem. Higher mass results in greater
// gravitational force according to Newton's law of gravitation.
//
// For attractors, mass typically ranges from 1000-100000.
// For particles, mass is usually 1.0 for uniform behavior.
//
// Example creating a heavy attractor:
//
//	mass := components.NewMass().WithValue(50000)
type Mass struct {
	// Value is the gravitational mass in arbitrary units.
	Value float32
}

// Mask returns the component mask for Mass.
func (m *Mass) Mask() uint64 { return MaskMass }

// NewMass creates a new Mass component with default value of 1.
func NewMass() *Mass { return &Mass{Value: 1.0} }

// WithValue sets the mass value and returns the mass for chaining.
func (m *Mass) WithValue(v float32) *Mass { m.Value = v; return m }
