package systems

import (
	"testing"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
)

// TestColorSystem tests color interpolation based on lifetime.
func TestColorSystem_Process(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	// Create a particle with color and lifetime
	entity := ecs.NewEntity("test-particle", []ecs.Component{
		components.NewColor().WithGradient(255, 0, 0, 255, 0, 0, 255, 0),
		components.NewLifetime().WithTTL(2.0),
		components.NewSize().WithRadius(10.0).WithEndSize(1.0),
	})
	em.Add(entity)

	// Process should not panic
	sys.Setup()
	result := sys.Process(em)
	sys.Teardown()

	if result != ecs.StateEngineContinue {
		t.Errorf("expected StateEngineContinue, got %d", result)
	}
}

// TestColorSystem_ColorInterpolation tests that colors interpolate correctly.
func TestColorSystem_ColorInterpolation(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	// Create particle at 50% lifetime
	color := components.NewColor().WithGradient(255, 0, 0, 255, 0, 0, 255, 0)
	lifetime := components.NewLifetime().WithTTL(2.0)
	lifetime.Age = 1.0 // 50% through

	entity := ecs.NewEntity("test", []ecs.Component{color, lifetime})
	em.Add(entity)

	sys.Process(em)

	// At 50%, color should be interpolated between start and end
	// Start: 255,0,0,255 End: 0,0,255,0
	// Expected at 50%: ~127,0,127,127
	if color.R > 200 || color.R < 50 {
		t.Errorf("expected R to be around 127, got %d", color.R)
	}
}

// TestColorSystem_SizeInterpolation tests that size interpolates correctly.
func TestColorSystem_SizeInterpolation(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	// Create particle at 50% lifetime
	size := components.NewSize().WithRadius(10.0).WithEndSize(2.0)
	lifetime := components.NewLifetime().WithTTL(2.0)
	lifetime.Age = 1.0 // 50% through

	entity := ecs.NewEntity("test", []ecs.Component{
		components.NewColor(),
		lifetime,
		size,
	})
	em.Add(entity)

	sys.Process(em)

	// At 50%, size should be around 6 (midpoint between 10 and 2)
	if size.Radius > 8 || size.Radius < 4 {
		t.Errorf("expected Radius to be around 6, got %f", size.Radius)
	}
}

// TestGravitySystem tests gravitational force calculation.
func TestGravitySystem_Process(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	// Create an attractor
	attractor := ecs.NewEntity("attractor", []ecs.Component{
		components.NewPosition().With(500, 500),
		components.NewMass().WithValue(10000),
		components.NewAttractor(),
	})
	em.Add(attractor)

	// Create a particle
	particle := ecs.NewEntity("particle", []ecs.Component{
		components.NewPosition().With(400, 400),
		components.NewAcceleration(),
		components.NewParticle(),
	})
	em.Add(particle)

	sys.Setup()
	result := sys.Process(em)
	sys.Teardown()

	if result != ecs.StateEngineContinue {
		t.Errorf("expected StateEngineContinue, got %d", result)
	}

	// Check that acceleration was applied (particle should be pulled toward attractor)
	acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
	if acc.X <= 0 || acc.Y <= 0 {
		t.Errorf("expected positive acceleration toward attractor, got (%f, %f)", acc.X, acc.Y)
	}
}

// TestGravitySystem_ZeroMass tests that zero mass attractors have no effect.
func TestGravitySystem_ZeroMass(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	// Create an attractor with zero mass
	attractor := ecs.NewEntity("attractor", []ecs.Component{
		components.NewPosition().With(500, 500),
		components.NewMass().WithValue(0),
		components.NewAttractor(),
	})
	em.Add(attractor)

	// Create a particle
	particle := ecs.NewEntity("particle", []ecs.Component{
		components.NewPosition().With(400, 400),
		components.NewAcceleration(),
		components.NewParticle(),
	})
	em.Add(particle)

	sys.Process(em)

	// With zero mass, no acceleration should be applied
	acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
	if acc.X != 0 || acc.Y != 0 {
		t.Errorf("expected zero acceleration with zero mass attractor, got (%f, %f)", acc.X, acc.Y)
	}
}

// TestGravitySystem_MultipleAttractors tests that multiple attractors sum forces.
func TestGravitySystem_MultipleAttractors(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	// Create two attractors on opposite sides
	attractor1 := ecs.NewEntity("attractor1", []ecs.Component{
		components.NewPosition().With(0, 250),
		components.NewMass().WithValue(5000),
		components.NewAttractor(),
	})
	em.Add(attractor1)

	attractor2 := ecs.NewEntity("attractor2", []ecs.Component{
		components.NewPosition().With(500, 250),
		components.NewMass().WithValue(5000),
		components.NewAttractor(),
	})
	em.Add(attractor2)

	// Create a particle in the middle
	particle := ecs.NewEntity("particle", []ecs.Component{
		components.NewPosition().With(250, 250),
		components.NewAcceleration(),
		components.NewParticle(),
	})
	em.Add(particle)

	sys.Process(em)

	// Forces should roughly cancel out (equal mass, equal distance)
	acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
	// Y should be ~0 since attractors are horizontally aligned with particle
	if acc.Y > 1 || acc.Y < -1 {
		t.Errorf("expected Y acceleration near 0, got %f", acc.Y)
	}
}

// TestLerp tests the linear interpolation function.
func TestLerp(t *testing.T) {
	tests := []struct {
		a, b     uint8
		tval     float32
		expected uint8
	}{
		{0, 100, 0.0, 0},
		{0, 100, 1.0, 100},
		{0, 100, 0.5, 50},
		{100, 0, 0.5, 50},
		{255, 0, 0.5, 127},
	}

	for _, tc := range tests {
		result := lerp(tc.a, tc.b, tc.tval)
		// Allow small rounding differences
		diff := int(result) - int(tc.expected)
		if diff > 1 || diff < -1 {
			t.Errorf("lerp(%d, %d, %f) = %d, want ~%d", tc.a, tc.b, tc.tval, result, tc.expected)
		}
	}
}

// TestLerpF tests the float linear interpolation function.
func TestLerpF(t *testing.T) {
	tests := []struct {
		a, b, tval float32
		expected   float32
	}{
		{0, 100, 0.0, 0},
		{0, 100, 1.0, 100},
		{0, 100, 0.5, 50},
		{100, 0, 0.5, 50},
		{10, 2, 0.5, 6},
	}

	for _, tc := range tests {
		result := lerpF(tc.a, tc.b, tc.tval)
		if result != tc.expected {
			t.Errorf("lerpF(%f, %f, %f) = %f, want %f", tc.a, tc.b, tc.tval, result, tc.expected)
		}
	}
}
