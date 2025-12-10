package components

// Mass represents the gravitational mass of an entity.
type Mass struct {
	Value float32
}

// Mask returns the component mask for Mass.
func (m *Mass) Mask() uint64 { return MaskMass }

// NewMass creates a new Mass component with default value of 1.
func NewMass() *Mass { return &Mass{Value: 1.0} }

// WithValue sets the mass value and returns the mass for chaining.
func (m *Mass) WithValue(v float32) *Mass { m.Value = v; return m }
