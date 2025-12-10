package components

// Particle is a tag component that identifies particle entities.
// Tag components have no data - they exist purely to categorize entities
// for efficient filtering in systems.
//
// Entities with the Particle tag are:
//   - Spawned by EmitterSystem
//   - Attracted by GravitySystem
//   - Rendered by RenderSystem
//   - Aged and removed by LifetimeSystem
//
// Example creating a particle entity:
//
//	entity := ecs.NewEntity("particle",
//	    components.NewParticle(),
//	    components.NewPosition(),
//	    components.NewVelocity(),
//	)
type Particle struct{}

// Mask returns the component mask for Particle.
func (p *Particle) Mask() uint64 { return MaskParticle }

// NewParticle creates a new Particle tag component.
func NewParticle() *Particle { return &Particle{} }
