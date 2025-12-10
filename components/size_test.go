package components

import (
	"testing"
)

func TestSize_Mask(t *testing.T) {
	s := NewSize()
	if s.Mask() != MaskSize {
		t.Errorf("Size.Mask() = %v, want %v", s.Mask(), MaskSize)
	}
}

func TestSize_NewSize(t *testing.T) {
	s := NewSize()
	if s.Radius != 3.0 {
		t.Errorf("NewSize().Radius = %v, want 3.0", s.Radius)
	}
	if s.StartSize != 3.0 {
		t.Errorf("NewSize().StartSize = %v, want 3.0", s.StartSize)
	}
	if s.EndSize != 1.0 {
		t.Errorf("NewSize().EndSize = %v, want 1.0", s.EndSize)
	}
}

func TestSize_WithRadius(t *testing.T) {
	s := NewSize().WithRadius(5.0)
	if s.Radius != 5.0 {
		t.Errorf("Size.WithRadius(5.0).Radius = %v, want 5.0", s.Radius)
	}
	if s.StartSize != 5.0 {
		t.Errorf("Size.WithRadius(5.0).StartSize = %v, want 5.0", s.StartSize)
	}
}

func TestSize_WithEndSize(t *testing.T) {
	s := NewSize().WithEndSize(0.5)
	if s.EndSize != 0.5 {
		t.Errorf("Size.WithEndSize(0.5).EndSize = %v, want 0.5", s.EndSize)
	}
}

func TestSize_Chaining(t *testing.T) {
	s := NewSize().WithRadius(10.0).WithEndSize(2.0)
	if s.Radius != 10.0 || s.StartSize != 10.0 || s.EndSize != 2.0 {
		t.Errorf("Size chaining failed: got Radius=%v, StartSize=%v, EndSize=%v, want 10.0, 10.0, 2.0",
			s.Radius, s.StartSize, s.EndSize)
	}
}
