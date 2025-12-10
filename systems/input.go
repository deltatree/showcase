package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// inputSystem handles mouse and keyboard input.
type inputSystem struct {
	mouseAttractorID string
	currentPreset    int
	presetSwitcher   func(int)
	lockedMode       int     // 0=none, 1=attract, -1=repel
	lastClickTime    float64 // for double-click detection
}

// NewInputSystem creates a new input system.
func NewInputSystem(presetSwitcher func(int)) ecs.System {
	return &inputSystem{
		mouseAttractorID: "mouse-attractor",
		currentPreset:    0,
		presetSwitcher:   presetSwitcher,
		lockedMode:       0,
		lastClickTime:    0,
	}
}

func (s *inputSystem) Setup() {}

func (s *inputSystem) Process(em ecs.EntityManager) (state int) {
	mouseX := float32(rl.GetMouseX())
	mouseY := float32(rl.GetMouseY())

	// Find or create mouse attractor
	attractor := em.Get(s.mouseAttractorID)
	if attractor == nil {
		attractor = ecs.NewEntity(s.mouseAttractorID, []ecs.Component{
			components.NewPosition().With(mouseX, mouseY),
			components.NewMass().WithValue(0),
			components.NewAttractor(),
		})
		em.Add(attractor)
	}

	// Update position
	pos := attractor.Get(components.MaskPosition).(*components.Position)
	pos.X, pos.Y = mouseX, mouseY

	// Update mass based on mouse buttons or locked mode
	mass := attractor.Get(components.MaskMass).(*components.Mass)
	currentTime := rl.GetTime()

	// Double-click detection for locking
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if currentTime-s.lastClickTime < 0.3 {
			// Double-click: toggle lock
			if s.lockedMode == 1 {
				s.lockedMode = 0
			} else {
				s.lockedMode = 1
			}
		}
		s.lastClickTime = currentTime
	}
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		if currentTime-s.lastClickTime < 0.3 {
			// Double-click: toggle lock
			if s.lockedMode == -1 {
				s.lockedMode = 0
			} else {
				s.lockedMode = -1
			}
		}
		s.lastClickTime = currentTime
	}

	// Apply mass based on locked mode or current button state
	if s.lockedMode == 1 {
		mass.Value = 8000 // Locked attract
	} else if s.lockedMode == -1 {
		mass.Value = -8000 // Locked repel
	} else if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mass.Value = 5000 // Attract (stronger)
	} else if rl.IsMouseButtonDown(rl.MouseRightButton) {
		mass.Value = -5000 // Repel (stronger)
	} else {
		mass.Value = 0 // Inactive
	}

	// Preset switching via keys 1-5
	if s.presetSwitcher != nil {
		if rl.IsKeyPressed(rl.KeyOne) {
			s.currentPreset = 0
			s.presetSwitcher(0)
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			s.currentPreset = 1
			s.presetSwitcher(1)
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			s.currentPreset = 2
			s.presetSwitcher(2)
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			s.currentPreset = 3
			s.presetSwitcher(3)
		}
		if rl.IsKeyPressed(rl.KeyFive) {
			s.currentPreset = 4
			s.presetSwitcher(4)
		}
	}

	return ecs.StateEngineContinue
}

func (s *inputSystem) Teardown() {}
