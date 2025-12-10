package premium

// QualityLevel represents the visual quality preset.
type QualityLevel int

const (
	QualityLow QualityLevel = iota
	QualityMedium
	QualityHigh
)

// String returns the quality level name.
func (q QualityLevel) String() string {
	switch q {
	case QualityLow:
		return "Low"
	case QualityMedium:
		return "Medium"
	case QualityHigh:
		return "High"
	default:
		return "Unknown"
	}
}

// QualitySettings contains all quality-dependent configuration.
type QualitySettings struct {
	Level        QualityLevel
	MaxParticles int
	GlowEnabled  bool
	GlowPasses   int
	MotionBlur   bool
	BlurSamples  int
	ParticleSize float32
	TrailLength  int
}

// Low quality preset - for older hardware
var qualityLow = QualitySettings{
	Level:        QualityLow,
	MaxParticles: 3000,
	GlowEnabled:  false,
	GlowPasses:   0,
	MotionBlur:   false,
	BlurSamples:  0,
	ParticleSize: 3.0,
	TrailLength:  0,
}

// Medium quality preset - balanced
var qualityMedium = QualitySettings{
	Level:        QualityMedium,
	MaxParticles: 7000,
	GlowEnabled:  true,
	GlowPasses:   1,
	MotionBlur:   false,
	BlurSamples:  0,
	ParticleSize: 2.5,
	TrailLength:  3,
}

// High quality preset - full visual experience
var qualityHigh = QualitySettings{
	Level:        QualityHigh,
	MaxParticles: 15000,
	GlowEnabled:  true,
	GlowPasses:   2,
	MotionBlur:   true,
	BlurSamples:  4,
	ParticleSize: 2.0,
	TrailLength:  5,
}

// GetQualitySettings returns settings for a quality level.
func GetQualitySettings(level QualityLevel) QualitySettings {
	switch level {
	case QualityLow:
		return qualityLow
	case QualityMedium:
		return qualityMedium
	case QualityHigh:
		return qualityHigh
	default:
		return qualityMedium
	}
}

// NextQuality cycles to the next quality level.
func NextQuality(current QualityLevel) QualityLevel {
	switch current {
	case QualityLow:
		return QualityMedium
	case QualityMedium:
		return QualityHigh
	default:
		return QualityLow
	}
}
