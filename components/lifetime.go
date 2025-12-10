package components

// Lifetime represents the lifespan of an entity, enabling automatic cleanup
// and time-based animations.
//
// The lifetime system increments Age each frame and marks entities as Expired
// when Age >= TTL, triggering their removal from the entity manager.
// The Progress() method returns the normalized age (0.0 to 1.0) for interpolation.
type Lifetime struct {
	// TTL (Time To Live) is the maximum lifespan in seconds.
	TTL float32
	// Age is the current age in seconds, incremented each frame.
	Age float32
	// Expired is set to true when the entity should be removed.
	Expired bool
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
