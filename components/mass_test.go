package components

import (
	"testing"
)

func TestMass_Mask(t *testing.T) {
	m := NewMass()
	if m.Mask() != MaskMass {
		t.Errorf("Mass.Mask() = %v, want %v", m.Mask(), MaskMass)
	}
}

func TestMass_NewMass(t *testing.T) {
	m := NewMass()
	if m.Value != 1.0 {
		t.Errorf("NewMass().Value = %v, want 1.0", m.Value)
	}
}

func TestMass_WithValue(t *testing.T) {
	m := NewMass().WithValue(500)
	if m.Value != 500 {
		t.Errorf("Mass.WithValue(500).Value = %v, want 500", m.Value)
	}
}

func TestMass_NegativeValue(t *testing.T) {
	m := NewMass().WithValue(-500)
	if m.Value != -500 {
		t.Errorf("Mass.WithValue(-500).Value = %v, want -500", m.Value)
	}
}

func TestMass_ZeroValue(t *testing.T) {
	m := NewMass().WithValue(0)
	if m.Value != 0 {
		t.Errorf("Mass.WithValue(0).Value = %v, want 0", m.Value)
	}
}
