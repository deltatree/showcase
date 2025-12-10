package components

// Emitter is a tag component to identify emitter entities.
type Emitter struct{}

// Mask returns the component mask for Emitter.
func (e *Emitter) Mask() uint64 { return MaskEmitter }

// NewEmitter creates a new Emitter tag component.
func NewEmitter() *Emitter { return &Emitter{} }
