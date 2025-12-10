package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestDefault(t *testing.T) {
	cfg := Default()

	if cfg.Window.Width != 1280 {
		t.Errorf("Default().Window.Width = %v, want 1280", cfg.Window.Width)
	}
	if cfg.Window.Height != 720 {
		t.Errorf("Default().Window.Height = %v, want 720", cfg.Window.Height)
	}
	if cfg.Window.Title != "Particle Symphony - ECS Showcase" {
		t.Errorf("Default().Window.Title = %v, want 'Particle Symphony - ECS Showcase'", cfg.Window.Title)
	}
	if cfg.Window.FPS != 60 {
		t.Errorf("Default().Window.FPS = %v, want 60", cfg.Window.FPS)
	}
	if cfg.Particles.MaxCount != 10000 {
		t.Errorf("Default().Particles.MaxCount = %v, want 10000", cfg.Particles.MaxCount)
	}
}

func TestLoad_DefaultOnMissingFile(t *testing.T) {
	cfg, err := Load("nonexistent_config.json")

	if err == nil {
		t.Error("Load() on missing file should return error")
	}

	if cfg.Window.Width != 1280 {
		t.Errorf("Load() on missing file: Window.Width = %v, want 1280", cfg.Window.Width)
	}
	if cfg.Window.Height != 720 {
		t.Errorf("Load() on missing file: Window.Height = %v, want 720", cfg.Window.Height)
	}
}

func TestLoad_ValidConfig(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "test_config.json")

	testCfg := Config{
		Window: WindowConfig{
			Width:  1920,
			Height: 1080,
			Title:  "Test Title",
			FPS:    120,
		},
		Particles: ParticleConfig{
			MaxCount:    50000,
			DefaultSize: 5.0,
			SpawnRate:   100,
			DefaultTTL:  3.0,
		},
		Physics: PhysicsConfig{
			Gravity:     500.0,
			Damping:     0.98,
			MaxVelocity: 1000.0,
		},
	}

	data, err := json.Marshal(testCfg)
	if err != nil {
		t.Fatalf("Failed to marshal test config: %v", err)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	cfg, err := Load(configPath)

	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.Window.Width != 1920 {
		t.Errorf("Load().Window.Width = %v, want 1920", cfg.Window.Width)
	}
	if cfg.Window.Height != 1080 {
		t.Errorf("Load().Window.Height = %v, want 1080", cfg.Window.Height)
	}
	if cfg.Window.Title != "Test Title" {
		t.Errorf("Load().Window.Title = %v, want 'Test Title'", cfg.Window.Title)
	}
	if cfg.Window.FPS != 120 {
		t.Errorf("Load().Window.FPS = %v, want 120", cfg.Window.FPS)
	}
	if cfg.Particles.MaxCount != 50000 {
		t.Errorf("Load().Particles.MaxCount = %v, want 50000", cfg.Particles.MaxCount)
	}
	if cfg.Physics.Gravity != 500.0 {
		t.Errorf("Load().Physics.Gravity = %v, want 500.0", cfg.Physics.Gravity)
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "invalid_config.json")

	err := os.WriteFile(configPath, []byte("not valid json"), 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid config: %v", err)
	}

	cfg, loadErr := Load(configPath)

	if loadErr == nil {
		t.Error("Load() on invalid JSON should return error")
	}

	if cfg.Window.Width != 1280 {
		t.Errorf("Load() on invalid JSON: Window.Width = %v, want 1280 (default)", cfg.Window.Width)
	}
}
