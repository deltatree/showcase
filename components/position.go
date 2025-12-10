package components

// Position represents the 2D coordinates of an entity.
type Position struct {
	X, Y float32
}

// Mask returns the component mask for Position.
func (p *Position) Mask() uint64 { return MaskPosition }

// NewPosition creates a new Position component.
func NewPosition() *Position { return &Position{} }

// WithX sets the X coordinate and returns the position for chaining.
func (p *Position) WithX(x float32) *Position { p.X = x; return p }

// WithY sets the Y coordinate and returns the position for chaining.
func (p *Position) WithY(y float32) *Position { p.Y = y; return p }

// With sets both X and Y coordinates and returns the position for chaining.
func (p *Position) With(x, y float32) *Position { p.X = x; p.Y = y; return p }
