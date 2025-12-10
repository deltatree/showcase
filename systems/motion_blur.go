package systems

import (
	"math"

	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/premium"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// MotionBlurRenderer provides motion blur rendering for fast-moving particles.
type MotionBlurRenderer struct {
	enabled   bool
	samples   int
	threshold float32 // Minimum speed for blur effect
}

// NewMotionBlurRenderer creates a new motion blur renderer.
func NewMotionBlurRenderer(enabled bool, samples int) *MotionBlurRenderer {
	return &MotionBlurRenderer{
		enabled:   enabled,
		samples:   samples,
		threshold: 100.0,
	}
}

// RenderWithBlur draws a particle with motion blur based on velocity.
func (r *MotionBlurRenderer) RenderWithBlur(
	pos *components.Position,
	vel *components.Velocity,
	col *components.Color,
	size *components.Size,
	shakeX, shakeY float32,
) {
	drawX := int32(pos.X + shakeX)
	drawY := int32(pos.Y + shakeY)
	drawColor := rl.NewColor(col.R, col.G, col.B, col.A)

	if !r.enabled || vel == nil {
		rl.DrawCircle(drawX, drawY, size.Radius, drawColor)
		return
	}

	// Calculate speed
	speed := float32(math.Sqrt(float64(vel.X*vel.X + vel.Y*vel.Y)))

	if speed < r.threshold {
		// Normal rendering for slow particles
		rl.DrawCircle(drawX, drawY, size.Radius, drawColor)
		return
	}

	// Motion blur: render trail of semi-transparent circles
	blurSteps := int(math.Min(float64(speed/50), float64(r.samples)))
	if blurSteps < 2 {
		blurSteps = 2
	}

	// Calculate step alpha
	stepAlpha := float32(col.A) / float32(blurSteps+1)

	// Time factor (~1 frame at 60fps)
	dt := float32(0.016)

	for i := blurSteps; i >= 1; i-- {
		t := float32(i) / float32(blurSteps)
		// Position back in time
		x := pos.X - vel.X*t*dt + shakeX
		y := pos.Y - vel.Y*t*dt + shakeY
		// Fade and shrink
		alpha := uint8(stepAlpha * (1.0 - t*0.5))
		trailSize := size.Radius * (1.0 - t*0.3)

		blurColor := rl.NewColor(col.R, col.G, col.B, alpha)
		rl.DrawCircle(int32(x), int32(y), trailSize, blurColor)
	}

	// Main particle on top
	rl.DrawCircle(drawX, drawY, size.Radius, drawColor)
}

// SetEnabled enables or disables motion blur.
func (r *MotionBlurRenderer) SetEnabled(enabled bool) {
	r.enabled = enabled
}

// SetSamples sets the number of blur samples.
func (r *MotionBlurRenderer) SetSamples(samples int) {
	r.samples = samples
}

// ApplyQuality applies quality settings to motion blur.
func (r *MotionBlurRenderer) ApplyQuality(q premium.QualitySettings) {
	r.enabled = q.MotionBlur
	r.samples = q.BlurSamples
}
