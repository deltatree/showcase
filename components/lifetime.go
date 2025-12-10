package components

// Lifetime represents the lifespan of an entity.
type Lifetime struct {
	TTL     float32 // Time To Live in seconds
	Age     float32 // Current age in seconds
	Expired bool    // Marked for removal
}

// Mask returns the component mask for Lifetime.
func (l *Lifetime) Mask() uint64 { return MaskLifetime }

// NewLifetime creates a new Lifetime component with default TTL of 5 seconds.
func NewLifetime() *Lifetime { return &Lifetime{TTL: 5.0} }

// WithTTL sets the Time To Live and returns the lifetime for chaining.
func (l *Lifetime) WithTTL(ttl float32) *Lifetime {
	l.TTL = ttl
	return l
}

// Progress returns the lifetime progress from 0.0 (new) to 1.0 (expired).
func (l *Lifetime) Progress() float32 {
	if l.TTL <= 0 {
		return 1.0
	}
	progress := l.Age / l.TTL
	if progress > 1.0 {
		return 1.0
	}
	return progress
}
