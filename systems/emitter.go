package systems

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/premium"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// EmitterSystem spawns new particles at configurable rates.
// It supports multiple spawn patterns and fully customizable particle
// properties including colors, sizes, velocities, and lifetimes.
//
// Spawn patterns:
//   - "random": particles spawn at random screen positions
//   - "center": particles spawn near screen center
//   - "emitter": particles spawn at emitter entity positions
type emitterSystem struct {
	spawnRate    int
	spawnTimer   float32
	maxParticles int
	width        float32
	height       float32
	rng          *rand.Rand
	idCounter    int64
	quality      premium.QualitySettings
	// Configurable spawn parameters set by presets
	StartColorR, StartColorG, StartColorB, StartColorA uint8
	EndColorR, EndColorG, EndColorB, EndColorA         uint8
	MinSize, MaxSize                                   float32
	MinTTL, MaxTTL                                     float32
	MinVel, MaxVel                                     float32
	SpawnPattern                                       string
}

// NewEmitterSystem creates a new emitter system with default parameters.
//
// Parameters:
//   - spawnRate: particles per second (0 disables spawning)
//   - maxParticles: maximum concurrent particles
//   - width, height: screen dimensions for spawn bounds
func NewEmitterSystem(spawnRate, maxParticles int, width, height float32) *emitterSystem {
	return &emitterSystem{
		spawnRate:    spawnRate,
		maxParticles: maxParticles,
		width:        width,
		height:       height,
		rng:          rand.New(rand.NewSource(time.Now().UnixNano())),
		quality:      premium.GetQualitySettings(premium.QualityMedium),
		StartColorR:  255, StartColorG: 150, StartColorB: 50, StartColorA: 255,
		EndColorR: 255, EndColorG: 50, EndColorB: 50, EndColorA: 0,
		MinSize:      2.0,
		MaxSize:      5.0,
		MinTTL:       3.0,
		MaxTTL:       5.0,
		MinVel:       -50,
		MaxVel:       50,
		SpawnPattern: "random",
	}
}

func (s *emitterSystem) Setup() {}

func (s *emitterSystem) Process(em ecs.EntityManager) (state int) {
	dt := rl.GetFrameTime()
	s.spawnTimer += dt

	particles := em.FilterByMask(components.MaskParticle)
	currentCount := len(particles)

	// Use quality-based max particles
	maxAllowed := s.quality.MaxParticles
	if s.maxParticles < maxAllowed {
		maxAllowed = s.maxParticles
	}

	if s.spawnRate <= 0 {
		return ecs.StateEngineContinue
	}
	spawnInterval := 1.0 / float32(s.spawnRate)

	for s.spawnTimer >= spawnInterval && currentCount < maxAllowed {
		s.spawnTimer -= spawnInterval
		s.spawnParticle(em)
		currentCount++
	}

	return ecs.StateEngineContinue
}

func (s *emitterSystem) spawnParticle(em ecs.EntityManager) {
	var x, y float32

	switch s.SpawnPattern {
	case "center":
		x = s.width/2 + (s.rng.Float32()-0.5)*100
		y = s.height/2 + (s.rng.Float32()-0.5)*100
	case "edges":
		side := s.rng.Intn(4)
		switch side {
		case 0:
			x = s.rng.Float32() * s.width
			y = 0
		case 1:
			x = s.rng.Float32() * s.width
			y = s.height
		case 2:
			x = 0
			y = s.rng.Float32() * s.height
		case 3:
			x = s.width
			y = s.rng.Float32() * s.height
		}
	default:
		x = s.rng.Float32() * s.width
		y = s.rng.Float32() * s.height
	}

	vx := s.MinVel + s.rng.Float32()*(s.MaxVel-s.MinVel)
	vy := s.MinVel + s.rng.Float32()*(s.MaxVel-s.MinVel)
	size := s.MinSize + s.rng.Float32()*(s.MaxSize-s.MinSize)
	ttl := s.MinTTL + s.rng.Float32()*(s.MaxTTL-s.MinTTL)

	s.idCounter++
	id := fmt.Sprintf("p-%d", s.idCounter)

	em.Add(ecs.NewEntity(id, []ecs.Component{
		components.NewPosition().With(x, y),
		components.NewVelocity().With(vx, vy),
		components.NewAcceleration(),
		components.NewColor().WithGradient(
			s.StartColorR, s.StartColorG, s.StartColorB, s.StartColorA,
			s.EndColorR, s.EndColorG, s.EndColorB, s.EndColorA,
		),
		components.NewLifetime().WithTTL(ttl),
		components.NewSize().WithRadius(size).WithEndSize(size * 0.3),
		components.NewParticle(),
	}))
}

func (s *emitterSystem) Teardown() {}

// SetColors sets the start and end colors for spawned particles.
func (s *emitterSystem) SetColors(sr, sg, sb, sa, er, eg, eb, ea uint8) {
	s.StartColorR, s.StartColorG, s.StartColorB, s.StartColorA = sr, sg, sb, sa
	s.EndColorR, s.EndColorG, s.EndColorB, s.EndColorA = er, eg, eb, ea
}

// SetSpawnPattern sets the spawn pattern for particles.
func (s *emitterSystem) SetSpawnPattern(pattern string) {
	s.SpawnPattern = pattern
}

// SetSpawnRate sets the spawn rate.
func (s *emitterSystem) SetSpawnRate(rate int) {
	s.spawnRate = rate
}

// SetMaxParticles sets the maximum particle count.
func (s *emitterSystem) SetMaxParticles(max int) {
	s.maxParticles = max
}

// GetMaxParticles returns the current max particle count.
func (s *emitterSystem) GetMaxParticles() int {
	return s.maxParticles
}

// SetQuality sets the quality level for particle limits.
func (s *emitterSystem) SetQuality(level premium.QualityLevel) {
	s.quality = premium.GetQualitySettings(level)
}

// GetQuality returns the current quality settings.
func (s *emitterSystem) GetQuality() premium.QualitySettings {
	return s.quality
}
