package components

import (
	"testing"
)

func TestParticle_Mask(t *testing.T) {
	p := NewParticle()
	if p.Mask() != MaskParticle {
		t.Errorf("Particle.Mask() = %v, want %v", p.Mask(), MaskParticle)
	}
}

func TestParticle_NewParticle(t *testing.T) {
	p := NewParticle()
	if p == nil {
		t.Error("NewParticle() returned nil")
	}
}

func TestAttractor_Mask(t *testing.T) {
	a := NewAttractor()
	if a.Mask() != MaskAttractor {
		t.Errorf("Attractor.Mask() = %v, want %v", a.Mask(), MaskAttractor)
	}
}

func TestAttractor_NewAttractor(t *testing.T) {
	a := NewAttractor()
	if a == nil {
		t.Error("NewAttractor() returned nil")
	}
}

func TestEmitter_Mask(t *testing.T) {
	e := NewEmitter()
	if e.Mask() != MaskEmitter {
		t.Errorf("Emitter.Mask() = %v, want %v", e.Mask(), MaskEmitter)
	}
}

func TestEmitter_NewEmitter(t *testing.T) {
	e := NewEmitter()
	if e == nil {
		t.Error("NewEmitter() returned nil")
	}
}
