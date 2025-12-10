package components

import (
	"testing"
)

func TestMaskPosition(t *testing.T) {
	if MaskPosition != uint64(1<<0) {
		t.Errorf("MaskPosition = %v, want %v", MaskPosition, uint64(1<<0))
	}
}

func TestMaskVelocity(t *testing.T) {
	if MaskVelocity != uint64(1<<1) {
		t.Errorf("MaskVelocity = %v, want %v", MaskVelocity, uint64(1<<1))
	}
}

func TestMaskAcceleration(t *testing.T) {
	if MaskAcceleration != uint64(1<<2) {
		t.Errorf("MaskAcceleration = %v, want %v", MaskAcceleration, uint64(1<<2))
	}
}

func TestMaskColor(t *testing.T) {
	if MaskColor != uint64(1<<3) {
		t.Errorf("MaskColor = %v, want %v", MaskColor, uint64(1<<3))
	}
}

func TestMaskLifetime(t *testing.T) {
	if MaskLifetime != uint64(1<<4) {
		t.Errorf("MaskLifetime = %v, want %v", MaskLifetime, uint64(1<<4))
	}
}

func TestMaskMass(t *testing.T) {
	if MaskMass != uint64(1<<5) {
		t.Errorf("MaskMass = %v, want %v", MaskMass, uint64(1<<5))
	}
}

func TestMaskSize(t *testing.T) {
	if MaskSize != uint64(1<<6) {
		t.Errorf("MaskSize = %v, want %v", MaskSize, uint64(1<<6))
	}
}

func TestMaskEmitter(t *testing.T) {
	if MaskEmitter != uint64(1<<7) {
		t.Errorf("MaskEmitter = %v, want %v", MaskEmitter, uint64(1<<7))
	}
}

func TestMaskAttractor(t *testing.T) {
	if MaskAttractor != uint64(1<<8) {
		t.Errorf("MaskAttractor = %v, want %v", MaskAttractor, uint64(1<<8))
	}
}

func TestMaskParticle(t *testing.T) {
	if MaskParticle != uint64(1<<9) {
		t.Errorf("MaskParticle = %v, want %v", MaskParticle, uint64(1<<9))
	}
}

func TestMaskMovable(t *testing.T) {
	expected := MaskPosition | MaskVelocity
	if MaskMovable != expected {
		t.Errorf("MaskMovable = %v, want %v", MaskMovable, expected)
	}
}

func TestMaskPhysics(t *testing.T) {
	expected := MaskMovable | MaskAcceleration
	if MaskPhysics != expected {
		t.Errorf("MaskPhysics = %v, want %v", MaskPhysics, expected)
	}
}

func TestMaskRenderable(t *testing.T) {
	expected := MaskPosition | MaskColor | MaskSize
	if MaskRenderable != expected {
		t.Errorf("MaskRenderable = %v, want %v", MaskRenderable, expected)
	}
}

func TestMaskFullParticle(t *testing.T) {
	expected := MaskPhysics | MaskColor | MaskLifetime | MaskSize | MaskParticle
	if MaskFullParticle != expected {
		t.Errorf("MaskFullParticle = %v, want %v", MaskFullParticle, expected)
	}
}

func TestNoMaskCollisions(t *testing.T) {
	masks := []uint64{
		MaskPosition,
		MaskVelocity,
		MaskAcceleration,
		MaskColor,
		MaskLifetime,
		MaskMass,
		MaskSize,
		MaskEmitter,
		MaskAttractor,
		MaskParticle,
	}

	for i := 0; i < len(masks); i++ {
		for j := i + 1; j < len(masks); j++ {
			if masks[i] == masks[j] {
				t.Errorf("Mask collision detected: masks[%d] == masks[%d] == %v", i, j, masks[i])
			}
			if masks[i]&masks[j] != 0 {
				t.Errorf("Mask overlap detected: masks[%d] & masks[%d] = %v", i, j, masks[i]&masks[j])
			}
		}
	}
}
