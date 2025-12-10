package components

// Attractor is a tag component to identify attractor entities.
type Attractor struct{}

// Mask returns the component mask for Attractor.
func (a *Attractor) Mask() uint64 { return MaskAttractor }

// NewAttractor creates a new Attractor tag component.
func NewAttractor() *Attractor { return &Attractor{} }
