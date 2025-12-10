package systems

import (
	"fmt"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// RenderSystem handles window management and rendering of all visible entities.
// It initializes the raylib window, handles window close events, and draws
// particles as filled circles with their current color and size.
//
// Debug overlay (toggle with F3) displays:
//   - FPS counter
//   - Active entity count
//   - Current preset name
//   - Mouse coordinates
//   - Control hints
type renderSystem struct {
	width, height int32
	title         string
	showDebug     bool
	presetName    string
}

// NewRenderSystem creates a new render system for the specified window.
//
// Parameters:
//   - width, height: window dimensions in pixels
//   - title: window title displayed in the title bar
func NewRenderSystem(width, height int32, title string) *renderSystem {
	return &renderSystem{
		width:      width,
		height:     height,
		title:      title,
		showDebug:  true,
		presetName: "Galaxy",
	}
}

func (s *renderSystem) Setup() {
	rl.InitWindow(s.width, s.height, s.title)
	rl.SetTargetFPS(60)
}

func (s *renderSystem) Process(em ecs.EntityManager) (state int) {
	if rl.WindowShouldClose() {
		return ecs.StateEngineStop
	}

	if rl.IsKeyPressed(rl.KeyF3) {
		s.showDebug = !s.showDebug
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(10, 10, 20, 255))

	particles := em.FilterByMask(components.MaskRenderable)
	for _, e := range particles {
		pos := e.Get(components.MaskPosition).(*components.Position)
		col := e.Get(components.MaskColor).(*components.Color)
		size := e.Get(components.MaskSize).(*components.Size)

		rl.DrawCircle(
			int32(pos.X),
			int32(pos.Y),
			size.Radius,
			rl.NewColor(col.R, col.G, col.B, col.A),
		)
	}

	if s.showDebug {
		rl.DrawFPS(10, 10)
		rl.DrawText(
			fmt.Sprintf("Entities: %d", len(particles)),
			10, 35, 20, rl.White,
		)
		rl.DrawText(
			fmt.Sprintf("Preset: %s", s.presetName),
			10, 60, 20, rl.White,
		)
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		rl.DrawText(
			fmt.Sprintf("Mouse: (%d, %d)", mouseX, mouseY),
			10, 85, 16, rl.Gray,
		)
		rl.DrawText(
			"F3: Toggle Debug | LMB: Attract | RMB: Repel | 1-5: Presets",
			10, s.height-30, 16, rl.Gray,
		)
	}

	rl.EndDrawing()

	return ecs.StateEngineContinue
}

func (s *renderSystem) Teardown() {
	rl.CloseWindow()
}

// SetPresetName sets the current preset name for display.
func (s *renderSystem) SetPresetName(name string) {
	s.presetName = name
}
