package components

// Attractor is a tag component that identifies attractor entities.
// Attractors exert gravitational forces on particles through GravitySystem,
// pulling or pushing particles based on Mass and distance.
//
// Multiple attractors can exist simultaneously, creating complex
// gravitational fields. Attractor positions can be animated for
// dynamic effects like orbiting gravity wells.
//
// An attractor entity needs:
//   - Attractor (tag)
//   - Position (gravity center)
//   - Mass (gravitational strength)
//
// Example creating an attractor:
//
//	entity := ecs.NewEntity("attractor",
//	    components.NewAttractor(),
//	    components.NewPosition().WithX(400).WithY(300),
//	    components.NewMass().WithValue(50000),
//	)
type Attractor struct{}

// Mask returns the component mask for Attractor.
func (a *Attractor) Mask() uint64 { return MaskAttractor }

// NewAttractor creates a new Attractor tag component.
func NewAttractor() *Attractor { return &Attractor{} }
