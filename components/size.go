package components

// Size represents the visual size of an entity.
type Size struct {
	Radius    float32
	StartSize float32
	EndSize   float32
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
