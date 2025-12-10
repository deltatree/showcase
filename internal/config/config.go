// Package config provides configuration loading and default values.
package config

import (
	"encoding/json"
	"os"
)

// Config holds all configuration for Particle Symphony.
type Config struct {
	Window    WindowConfig   `json:"window"`
	Particles ParticleConfig `json:"particles"`
	Physics   PhysicsConfig  `json:"physics"`
}

// WindowConfig holds window-related settings.
type WindowConfig struct {
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Title  string `json:"title"`
	FPS    int32  `json:"fps"`
}

// ParticleConfig holds particle-related settings.
type ParticleConfig struct {
	MaxCount    int     `json:"maxCount"`
	DefaultSize float32 `json:"defaultSize"`
	SpawnRate   int     `json:"spawnRate"`
	DefaultTTL  float32 `json:"defaultTTL"`
}

// PhysicsConfig holds physics-related settings.
type PhysicsConfig struct {
	Gravity     float32 `json:"gravity"`
	Damping     float32 `json:"damping"`
	MaxVelocity float32 `json:"maxVelocity"`
}

// Default returns sensible default configuration.
func Default() *Config {
	return &Config{
		Window: WindowConfig{
			Width:  1280,
			Height: 720,
			Title:  "Particle Symphony - ECS Showcase",
			FPS:    60,
		},
		Particles: ParticleConfig{
			MaxCount:    10000,
			DefaultSize: 3.0,
			SpawnRate:   100,
			DefaultTTL:  5.0,
		},
		Physics: PhysicsConfig{
			Gravity:     0.0,
			Damping:     0.99,
			MaxVelocity: 500.0,
		},
	}
}

// Load reads configuration from a JSON file. Returns default config on error.
func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Default(), err
	}

	cfg := Default()
	if err := json.Unmarshal(data, cfg); err != nil {
		return Default(), err
	}

	return cfg, nil
}
