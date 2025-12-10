package components

// Particle is a tag component to identify particle entities.
type Particle struct{}

// Mask returns the component mask for Particle.
func (p *Particle) Mask() uint64 { return MaskParticle }

// NewParticle creates a new Particle tag component.
func NewParticle() *Particle { return &Particle{} }
