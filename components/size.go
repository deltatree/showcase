package components

// Size represents the visual size of a particle entity.
// It supports animated size transitions from StartSize to EndSize
// over the particle's lifetime, enabling effects like fading sparks.
//
// Size interpolation is performed by ColorSystem based on the
// particle's age relative to its lifetime (calculated as t = Age/TTL).
//
// Example creating a particle that shrinks:
//
//	size := components.NewSize().WithRadius(5.0).WithEndSize(0.5)
type Size struct {
	// Radius is the current visual radius in pixels.
	Radius float32
	// StartSize is the initial radius when the particle spawns.
	StartSize float32
	// EndSize is the target radius when the particle expires.
	EndSize float32
}

// Mask returns the component mask for Size.
func (s *Size) Mask() uint64 { return MaskSize }

// NewSize creates a new Size component with default radius of 3.
func NewSize() *Size {
	return &Size{Radius: 3.0, StartSize: 3.0, EndSize: 1.0}
}

// WithRadius sets the radius and start size.
func (s *Size) WithRadius(r float32) *Size {
	s.Radius = r
	s.StartSize = r
	return s
}

// WithEndSize sets the end size for interpolation.
func (s *Size) WithEndSize(r float32) *Size {
	s.EndSize = r
	return s
}
