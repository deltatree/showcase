package components

import (
	"math"
	"testing"
)

func TestLifetime_Mask(t *testing.T) {
	l := NewLifetime()
	if l.Mask() != MaskLifetime {
		t.Errorf("Lifetime.Mask() = %v, want %v", l.Mask(), MaskLifetime)
	}
}

func TestLifetime_NewLifetime(t *testing.T) {
	l := NewLifetime()
	if l.TTL != 5.0 {
		t.Errorf("NewLifetime().TTL = %v, want 5.0", l.TTL)
	}
	if l.Age != 0 {
		t.Errorf("NewLifetime().Age = %v, want 0", l.Age)
	}
	if l.Expired {
		t.Error("NewLifetime().Expired = true, want false")
	}
}

func TestLifetime_WithTTL(t *testing.T) {
	l := NewLifetime().WithTTL(10.0)
	if l.TTL != 10.0 {
		t.Errorf("Lifetime.WithTTL(10.0).TTL = %v, want 10.0", l.TTL)
	}
}

func TestLifetime_Progress(t *testing.T) {
	tests := []struct {
		name string
		ttl  float32
		age  float32
		want float32
	}{
		{"start", 5.0, 0.0, 0.0},
		{"half", 5.0, 2.5, 0.5},
		{"end", 5.0, 5.0, 1.0},
		{"over", 5.0, 10.0, 1.0},
		{"zero ttl", 0.0, 1.0, 1.0},
		{"negative ttl", -1.0, 0.0, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLifetime().WithTTL(tt.ttl)
			l.Age = tt.age
			got := l.Progress()
			if math.Abs(float64(got-tt.want)) > 0.001 {
				t.Errorf("Progress() = %v, want %v", got, tt.want)
			}
		})
	}
}
