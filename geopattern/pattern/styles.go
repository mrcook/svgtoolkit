package pattern

// StylePreset represents a generator style.
type StylePreset struct {
	FillColourDark  string
	FillColourLight string
	StrokeColour    string
	StrokeOpacity   float64
	OpacityMin      float64
	OpacityMax      float64
}

// DefaultStylePresets as used by the pattern generators.
func DefaultStylePresets() *StylePreset {
	return &StylePreset{
		FillColourDark:  "#222",
		FillColourLight: "#ddd",
		StrokeColour:    "#000",
		StrokeOpacity:   0.02,
		OpacityMin:      0.02,
		OpacityMax:      0.15,
	}
}
