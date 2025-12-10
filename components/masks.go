// Package components contains all ECS component definitions for Particle Symphony.
package components

// Component masks for efficient entity filtering using bitmasks.
const (
	MaskPosition     = uint64(1 << 0)
	MaskVelocity     = uint64(1 << 1)
	MaskAcceleration = uint64(1 << 2)
	MaskColor        = uint64(1 << 3)
	MaskLifetime     = uint64(1 << 4)
	MaskMass         = uint64(1 << 5)
	MaskSize         = uint64(1 << 6)
	MaskEmitter      = uint64(1 << 7)
	MaskAttractor    = uint64(1 << 8)
	MaskParticle     = uint64(1 << 9)
)

// Composite masks for common component combinations.
const (
	MaskMovable      = MaskPosition | MaskVelocity
	MaskPhysics      = MaskMovable | MaskAcceleration
	MaskRenderable   = MaskPosition | MaskColor | MaskSize
	MaskFullParticle = MaskPhysics | MaskColor | MaskLifetime | MaskSize | MaskParticle
)
