package systems

import (
	"testing"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/premium"
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

// TestColorSystem_NoEntities tests ColorSystem with empty entity manager.
func TestColorSystem_NoEntities(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	result := sys.Process(em)
	if result != ecs.StateEngineContinue {
		t.Errorf("expected StateEngineContinue, got %d", result)
	}
}

// TestColorSystem_FullInterpolation tests color at various progress points.
func TestColorSystem_FullInterpolation(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	// Test at 0% (start)
	color := components.NewColor().WithGradient(255, 0, 0, 255, 0, 255, 0, 0)
	lifetime := components.NewLifetime().WithTTL(2.0)
	lifetime.Age = 0 // 0% through

	entity := ecs.NewEntity("test", []ecs.Component{color, lifetime})
	em.Add(entity)

	sys.Process(em)

	if color.R != 255 {
		t.Errorf("at 0%%, expected R=255, got %d", color.R)
	}
	if color.G != 0 {
		t.Errorf("at 0%%, expected G=0, got %d", color.G)
	}
}

// TestColorSystem_EndInterpolation tests color at 100% progress.
func TestColorSystem_EndInterpolation(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	color := components.NewColor().WithGradient(255, 0, 0, 255, 0, 255, 0, 0)
	lifetime := components.NewLifetime().WithTTL(2.0)
	lifetime.Age = 2.0 // 100% through

	entity := ecs.NewEntity("test", []ecs.Component{color, lifetime})
	em.Add(entity)

	sys.Process(em)

	if color.R != 0 {
		t.Errorf("at 100%%, expected R=0, got %d", color.R)
	}
	if color.G != 255 {
		t.Errorf("at 100%%, expected G=255, got %d", color.G)
	}
}

// TestGravitySystem_NegativeMass tests repulsion with negative mass.
func TestGravitySystem_NegativeMass(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	// Create an attractor with negative mass (repulsion)
	attractor := ecs.NewEntity("attractor", []ecs.Component{
		components.NewPosition().With(500, 500),
		components.NewMass().WithValue(-10000),
		components.NewAttractor(),
	})
	em.Add(attractor)

	// Create a particle to the left and above
	particle := ecs.NewEntity("particle", []ecs.Component{
		components.NewPosition().With(400, 400),
		components.NewAcceleration(),
		components.NewParticle(),
	})
	em.Add(particle)

	sys.Process(em)

	// With negative mass, particle should be pushed away (negative acceleration)
	acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
	if acc.X >= 0 || acc.Y >= 0 {
		t.Errorf("expected negative acceleration (repulsion), got (%f, %f)", acc.X, acc.Y)
	}
}

// TestGravitySystem_MinimumDistance tests that minimum distance is enforced.
func TestGravitySystem_MinimumDistance(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	// Create attractor and particle at same position
	attractor := ecs.NewEntity("attractor", []ecs.Component{
		components.NewPosition().With(100, 100),
		components.NewMass().WithValue(10000),
		components.NewAttractor(),
	})
	em.Add(attractor)

	particle := ecs.NewEntity("particle", []ecs.Component{
		components.NewPosition().With(100, 100),
		components.NewAcceleration(),
		components.NewParticle(),
	})
	em.Add(particle)

	// Should not panic (minimum distance prevents division by zero)
	sys.Process(em)

	// Acceleration should be finite (not NaN or Inf)
	acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
	if acc.X != acc.X || acc.Y != acc.Y { // NaN check
		t.Error("acceleration is NaN")
	}
}

// TestGravitySystem_NoAttractors tests gravity with no attractors present.
func TestGravitySystem_NoAttractors(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	particle := ecs.NewEntity("particle", []ecs.Component{
		components.NewPosition().With(100, 100),
		components.NewAcceleration().WithX(50).WithY(50),
		components.NewParticle(),
	})
	em.Add(particle)

	sys.Process(em)

	// Acceleration should be reset to zero
	acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
	if acc.X != 0 || acc.Y != 0 {
		t.Errorf("expected zero acceleration with no attractors, got (%f, %f)", acc.X, acc.Y)
	}
}

// TestGravitySystem_NoParticles tests gravity with no particles present.
func TestGravitySystem_NoParticles(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	attractor := ecs.NewEntity("attractor", []ecs.Component{
		components.NewPosition().With(500, 500),
		components.NewMass().WithValue(10000),
		components.NewAttractor(),
	})
	em.Add(attractor)

	result := sys.Process(em)
	if result != ecs.StateEngineContinue {
		t.Errorf("expected StateEngineContinue, got %d", result)
	}
}

// TestColorSystem_SizeOnly tests size interpolation without color changes.
func TestColorSystem_SizeOnly(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	size := components.NewSize().WithRadius(20.0).WithEndSize(5.0)
	lifetime := components.NewLifetime().WithTTL(4.0)
	lifetime.Age = 2.0 // 50%

	entity := ecs.NewEntity("test", []ecs.Component{
		components.NewColor(),
		lifetime,
		size,
	})
	em.Add(entity)

	sys.Process(em)

	// At 50%, size should be around 12.5 (midpoint between 20 and 5)
	expectedSize := 20.0 + (5.0-20.0)*0.5
	if size.Radius != float32(expectedSize) {
		t.Errorf("expected Radius %f, got %f", expectedSize, size.Radius)
	}
}

// TestLerp_EdgeCases tests edge cases for lerp function.
func TestLerp_EdgeCases(t *testing.T) {
	// Test same values
	result := lerp(100, 100, 0.5)
	if result != 100 {
		t.Errorf("lerp(100, 100, 0.5) = %d, want 100", result)
	}

	// Test zero to max
	result = lerp(0, 255, 1.0)
	if result != 255 {
		t.Errorf("lerp(0, 255, 1.0) = %d, want 255", result)
	}

	// Test max to zero
	result = lerp(255, 0, 1.0)
	if result != 0 {
		t.Errorf("lerp(255, 0, 1.0) = %d, want 0", result)
	}
}

// TestLerpF_EdgeCases tests edge cases for lerpF function.
func TestLerpF_EdgeCases(t *testing.T) {
	// Test same values
	result := lerpF(50.0, 50.0, 0.5)
	if result != 50.0 {
		t.Errorf("lerpF(50, 50, 0.5) = %f, want 50", result)
	}

	// Test negative values
	result = lerpF(-10.0, 10.0, 0.5)
	if result != 0.0 {
		t.Errorf("lerpF(-10, 10, 0.5) = %f, want 0", result)
	}

	// Test t out of range (should still work mathematically)
	result = lerpF(0, 100, 2.0)
	if result != 200.0 {
		t.Errorf("lerpF(0, 100, 2.0) = %f, want 200", result)
	}
}

// TestGravitySystem_LargeDistance tests gravity at large distances.
func TestGravitySystem_LargeDistance(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewGravitySystem()

	attractor := ecs.NewEntity("attractor", []ecs.Component{
		components.NewPosition().With(0, 0),
		components.NewMass().WithValue(10000),
		components.NewAttractor(),
	})
	em.Add(attractor)

	particle := ecs.NewEntity("particle", []ecs.Component{
		components.NewPosition().With(10000, 10000),
		components.NewAcceleration(),
		components.NewParticle(),
	})
	em.Add(particle)

	sys.Process(em)

	// At large distance, acceleration should be very small
	acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
	if acc.X > 1 || acc.Y > 1 {
		t.Errorf("expected very small acceleration at large distance, got (%f, %f)", acc.X, acc.Y)
	}
}

// TestColorSystem_MultipleEntities tests processing multiple entities.
func TestColorSystem_MultipleEntities(t *testing.T) {
	em := ecs.NewEntityManager()
	sys := NewColorSystem()

	// Create multiple particles at different progress levels
	for i := 0; i < 10; i++ {
		color := components.NewColor().WithGradient(255, 0, 0, 255, 0, 0, 255, 0)
		lifetime := components.NewLifetime().WithTTL(10.0)
		lifetime.Age = float32(i) // 0% to 90%

		entity := ecs.NewEntity("test-"+string(rune('0'+i)), []ecs.Component{color, lifetime})
		em.Add(entity)
	}

	result := sys.Process(em)
	if result != ecs.StateEngineContinue {
		t.Errorf("expected StateEngineContinue, got %d", result)
	}
}

// TestGravitySystem_Setup tests the Setup method.
func TestGravitySystem_Setup(t *testing.T) {
	sys := NewGravitySystem()
	// Setup should not panic
	sys.Setup()
}

// TestGravitySystem_Teardown tests the Teardown method.
func TestGravitySystem_Teardown(t *testing.T) {
	sys := NewGravitySystem()
	// Teardown should not panic
	sys.Teardown()
}

// TestColorSystem_Setup tests the Setup method.
func TestColorSystem_Setup(t *testing.T) {
	sys := NewColorSystem()
	// Setup should not panic
	sys.Setup()
}

// TestColorSystem_Teardown tests the Teardown method.
func TestColorSystem_Teardown(t *testing.T) {
	sys := NewColorSystem()
	// Teardown should not panic
	sys.Teardown()
}

// TestNewColorSystem tests the constructor.
func TestNewColorSystem(t *testing.T) {
	sys := NewColorSystem()
	if sys == nil {
		t.Error("NewColorSystem returned nil")
	}
}

// TestNewGravitySystem tests the constructor.
func TestNewGravitySystem(t *testing.T) {
	sys := NewGravitySystem()
	if sys == nil {
		t.Error("NewGravitySystem returned nil")
	}
}

// TestNewLifetimeSystem tests the constructor.
func TestNewLifetimeSystem(t *testing.T) {
	sys := NewLifetimeSystem()
	if sys == nil {
		t.Error("NewLifetimeSystem returned nil")
	}
}

// TestLifetimeSystem_Setup tests the Setup method.
func TestLifetimeSystem_Setup(t *testing.T) {
	sys := NewLifetimeSystem()
	// Setup should not panic
	sys.Setup()
}

// TestLifetimeSystem_Teardown tests the Teardown method.
func TestLifetimeSystem_Teardown(t *testing.T) {
	sys := NewLifetimeSystem()
	// Teardown should not panic
	sys.Teardown()
}

// TestNewPhysicsSystem tests the constructor.
func TestNewPhysicsSystem(t *testing.T) {
	sys := NewPhysicsSystem(0.99, 500, 1280, 720)
	if sys == nil {
		t.Error("NewPhysicsSystem returned nil")
	}
}

// TestPhysicsSystem_Setup tests the Setup method.
func TestPhysicsSystem_Setup(t *testing.T) {
	sys := NewPhysicsSystem(0.99, 500, 1280, 720)
	// Setup should not panic
	sys.Setup()
}

// TestPhysicsSystem_Teardown tests the Teardown method.
func TestPhysicsSystem_Teardown(t *testing.T) {
	sys := NewPhysicsSystem(0.99, 500, 1280, 720)
	// Teardown should not panic
	sys.Teardown()
}

// TestNewEmitterSystem tests the constructor and default values.
func TestNewEmitterSystem(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	if sys == nil {
		t.Error("NewEmitterSystem returned nil")
	}
}

// TestEmitterSystem_Setup tests the Setup method.
func TestEmitterSystem_Setup(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	// Setup should not panic
	sys.Setup()
}

// TestEmitterSystem_Teardown tests the Teardown method.
func TestEmitterSystem_Teardown(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	// Teardown should not panic
	sys.Teardown()
}

// TestEmitterSystem_SetColors tests the SetColors method.
func TestEmitterSystem_SetColors(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	sys.SetColors(255, 128, 64, 255, 0, 0, 0, 0)

	if sys.StartColorR != 255 || sys.StartColorG != 128 || sys.StartColorB != 64 {
		t.Errorf("SetColors did not set start colors correctly")
	}
	if sys.EndColorR != 0 || sys.EndColorG != 0 || sys.EndColorB != 0 {
		t.Errorf("SetColors did not set end colors correctly")
	}
}

// TestEmitterSystem_SetSpawnPattern tests the SetSpawnPattern method.
func TestEmitterSystem_SetSpawnPattern(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	sys.SetSpawnPattern("center")

	if sys.SpawnPattern != "center" {
		t.Errorf("SetSpawnPattern did not set pattern, expected 'center', got '%s'", sys.SpawnPattern)
	}
}

// TestEmitterSystem_SetSpawnRate tests the SetSpawnRate method.
func TestEmitterSystem_SetSpawnRate(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	sys.SetSpawnRate(200)
	// Cannot directly verify due to private field, but method should not panic
}

// TestNewRenderSystem tests the constructor.
func TestNewRenderSystem(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")
	if sys == nil {
		t.Error("NewRenderSystem returned nil")
	}
}

// TestRenderSystem_SetPresetName tests the SetPresetName method.
func TestRenderSystem_SetPresetName(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")
	sys.SetPresetName("Firework")

	if sys.presetName != "Firework" {
		t.Errorf("SetPresetName did not set name, expected 'Firework', got '%s'", sys.presetName)
	}
}

// TestNewInputSystem tests the constructor.
func TestNewInputSystem(t *testing.T) {
	callback := func(preset int) {}
	sys := NewInputSystem(callback)
	if sys == nil {
		t.Error("NewInputSystem returned nil")
	}
}

// TestInputSystem_Setup tests the Setup method.
func TestInputSystem_Setup(t *testing.T) {
	sys := NewInputSystem(nil)
	// Setup should not panic
	sys.Setup()
}

// TestInputSystem_Teardown tests the Teardown method.
func TestInputSystem_Teardown(t *testing.T) {
	sys := NewInputSystem(nil)
	// Teardown should not panic
	sys.Teardown()
}

// --- Premium Integration Tests ---

// TestRenderSystem_PremiumFeatures tests premium feature integration.
func TestRenderSystem_PremiumFeatures(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")

	// Test quality setting
	sys.SetQuality(premium.QualityHigh)
	if sys.GetQuality().Level != premium.QualityHigh {
		t.Error("SetQuality did not set quality level")
	}

	// Test palette setting
	sys.SetPalette(premium.FireworkPalette)
	if sys.palette.Name != "Firework" {
		t.Error("SetPalette did not set palette")
	}

	// Test effects
	sys.ApplyShake(5.0, 0.5)
	if !sys.effects.IsActive() {
		t.Error("ApplyShake should activate effects")
	}

	sys.ApplyPulse(1.05, 0.3)
	// Effects should still be active
	if !sys.effects.IsActive() {
		t.Error("ApplyPulse should activate effects")
	}
}

// TestEmitterSystem_QualityIntegration tests quality-based particle limits.
func TestEmitterSystem_QualityIntegration(t *testing.T) {
	sys := NewEmitterSystem(100, 10000, 1280, 720)

	// Default should be Medium
	if sys.GetQuality().Level != premium.QualityMedium {
		t.Error("Default quality should be Medium")
	}

	// Set to Low
	sys.SetQuality(premium.QualityLow)
	if sys.GetQuality().MaxParticles != 3000 {
		t.Errorf("Low quality MaxParticles = %d, want 3000", sys.GetQuality().MaxParticles)
	}

	// Set to High
	sys.SetQuality(premium.QualityHigh)
	if sys.GetQuality().MaxParticles != 15000 {
		t.Errorf("High quality MaxParticles = %d, want 15000", sys.GetQuality().MaxParticles)
	}
}

// TestNewMotionBlurRenderer tests motion blur renderer creation.
func TestNewMotionBlurRenderer(t *testing.T) {
	r := NewMotionBlurRenderer(true, 4)
	if r == nil {
		t.Error("NewMotionBlurRenderer returned nil")
	}
	if !r.enabled {
		t.Error("enabled should be true")
	}
	if r.samples != 4 {
		t.Errorf("samples = %d, want 4", r.samples)
	}
}

// TestMotionBlurRenderer_SetEnabled tests enable/disable.
func TestMotionBlurRenderer_SetEnabled(t *testing.T) {
	r := NewMotionBlurRenderer(true, 4)
	r.SetEnabled(false)
	if r.enabled {
		t.Error("SetEnabled(false) did not disable")
	}
	r.SetEnabled(true)
	if !r.enabled {
		t.Error("SetEnabled(true) did not enable")
	}
}

// TestMotionBlurRenderer_ApplyQuality tests quality application.
func TestMotionBlurRenderer_ApplyQuality(t *testing.T) {
	r := NewMotionBlurRenderer(false, 0)

	high := premium.GetQualitySettings(premium.QualityHigh)
	r.ApplyQuality(high)

	if !r.enabled {
		t.Error("High quality should enable motion blur")
	}
	if r.samples != high.BlurSamples {
		t.Errorf("samples = %d, want %d", r.samples, high.BlurSamples)
	}
}

// TestNewGlowRenderer tests glow renderer creation.
func TestNewGlowRenderer(t *testing.T) {
	r := NewGlowRenderer(true, 2)
	if r == nil {
		t.Error("NewGlowRenderer returned nil")
	}
	if !r.enabled {
		t.Error("enabled should be true")
	}
	if r.passes != 2 {
		t.Errorf("passes = %d, want 2", r.passes)
	}
}

// TestGlowRenderer_SetPalette tests palette setting.
func TestGlowRenderer_SetPalette(t *testing.T) {
	r := NewGlowRenderer(true, 2)
	r.SetPalette(premium.ChaosPalette)

	if r.palette.Name != "Chaos" {
		t.Errorf("palette.Name = %s, want Chaos", r.palette.Name)
	}
}

// TestGlowRenderer_ApplyQuality tests quality application.
func TestGlowRenderer_ApplyQuality(t *testing.T) {
	r := NewGlowRenderer(true, 5)

	low := premium.GetQualitySettings(premium.QualityLow)
	r.ApplyQuality(low)

	if r.enabled {
		t.Error("Low quality should disable glow")
	}
	if r.passes != low.GlowPasses {
		t.Errorf("passes = %d, want %d", r.passes, low.GlowPasses)
	}
}

// TestGlowRenderer_IsEnabled tests IsEnabled method.
func TestGlowRenderer_IsEnabled(t *testing.T) {
	r := NewGlowRenderer(true, 2)
	if !r.IsEnabled() {
		t.Error("IsEnabled should return true")
	}
	r.SetEnabled(false)
	if r.IsEnabled() {
		t.Error("IsEnabled should return false after SetEnabled(false)")
	}
}

// TestGlowRenderer_SetPasses tests SetPasses method.
func TestGlowRenderer_SetPasses(t *testing.T) {
	r := NewGlowRenderer(true, 2)
	r.SetPasses(5)
	if r.passes != 5 {
		t.Errorf("SetPasses did not set passes, expected 5, got %d", r.passes)
	}
}

// TestMotionBlurRenderer_SetSamples tests SetSamples method.
func TestMotionBlurRenderer_SetSamples(t *testing.T) {
	r := NewMotionBlurRenderer(true, 4)
	r.SetSamples(8)
	if r.samples != 8 {
		t.Errorf("SetSamples did not set samples, expected 8, got %d", r.samples)
	}
}

// TestEmitterSystem_SetMaxParticles tests the SetMaxParticles method.
func TestEmitterSystem_SetMaxParticles(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	sys.SetMaxParticles(8000)
	if sys.GetMaxParticles() != 8000 {
		t.Errorf("SetMaxParticles did not set max, expected 8000, got %d", sys.GetMaxParticles())
	}
}

// TestEmitterSystem_GetMaxParticles tests the GetMaxParticles method.
func TestEmitterSystem_GetMaxParticles(t *testing.T) {
	sys := NewEmitterSystem(100, 5000, 1280, 720)
	if sys.GetMaxParticles() != 5000 {
		t.Errorf("GetMaxParticles wrong, expected 5000, got %d", sys.GetMaxParticles())
	}
}

// TestRenderSystem_SetMaxParticles tests the SetMaxParticles method.
func TestRenderSystem_SetMaxParticles(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")
	sys.SetMaxParticles(12000)
	if sys.GetMaxParticles() != 12000 {
		t.Errorf("SetMaxParticles did not set max, expected 12000, got %d", sys.GetMaxParticles())
	}
}

// TestRenderSystem_GetMaxParticles tests the GetMaxParticles method.
func TestRenderSystem_GetMaxParticles(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")
	// Default should be based on medium quality
	max := sys.GetMaxParticles()
	if max == 0 {
		t.Error("GetMaxParticles should not return 0")
	}
}

// TestRenderSystem_SetOnParticleChange tests the SetOnParticleChange callback.
func TestRenderSystem_SetOnParticleChange(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")
	sys.SetOnParticleChange(func(count int) {
		// Callback logic
	})
	// Callback is set but not invoked in this test (requires Process)
	if sys.onParticleChange == nil {
		t.Error("SetOnParticleChange did not set callback")
	}
}

// TestRenderSystem_FullscreenState tests the fullscreen state initialization.
func TestRenderSystem_FullscreenState(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")
	// Should start in windowed mode
	if sys.isFullscreen {
		t.Error("RenderSystem should start in windowed mode, not fullscreen")
	}
}

// TestRenderSystem_DefaultQuality tests that default quality is Medium for performance.
func TestRenderSystem_DefaultQuality(t *testing.T) {
	sys := NewRenderSystem(1280, 720, "Test")
	if sys.quality.Level != premium.QualityMedium {
		t.Errorf("Default quality should be Medium, got %s", sys.quality.Level.String())
	}
}
