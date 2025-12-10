package components

// Emitter is a tag component that identifies emitter entities.
// Emitters are the source of particles - they spawn new particles
// at their Position at a rate defined by the current preset.
//
// Emitter position is typically updated by InputSystem to follow
// the mouse cursor, creating interactive particle effects.
//
// An emitter entity needs:
//   - Emitter (tag)
//   - Position (spawn location)
//
// Example creating an emitter:
//
//	entity := ecs.NewEntity("emitter",
//	    components.NewEmitter(),
//	    components.NewPosition().WithX(640).WithY(360),
//	)
type Emitter struct{}

// Mask returns the component mask for Emitter.
func (e *Emitter) Mask() uint64 { return MaskEmitter }

// NewEmitter creates a new Emitter tag component.
func NewEmitter() *Emitter { return &Emitter{} }
