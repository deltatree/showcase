package components

// Acceleration represents the acceleration vector of an entity.
type Acceleration struct {
	X, Y float32
}

// Mask returns the component mask for Acceleration.
func (a *Acceleration) Mask() uint64 { return MaskAcceleration }

// NewAcceleration creates a new Acceleration component.
func NewAcceleration() *Acceleration { return &Acceleration{} }

// Reset sets both X and Y acceleration to zero.
func (a *Acceleration) Reset() { a.X = 0; a.Y = 0 }

// Add adds the given values to the current acceleration.
func (a *Acceleration) Add(x, y float32) { a.X += x; a.Y += y }

// WithX sets the X acceleration and returns the acceleration for chaining.
func (a *Acceleration) WithX(x float32) *Acceleration { a.X = x; return a }

// WithY sets the Y acceleration and returns the acceleration for chaining.
func (a *Acceleration) WithY(y float32) *Acceleration { a.Y = y; return a }
