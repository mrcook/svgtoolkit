// Package pattern is used to generate various geo patterns.
//
//		pat := pattern.New(400, 400, "The Pattern Package", canvas)
//		pat.Generate()
//
// Then add additional elements to the svg.Canvas as needed.
package pattern

import (
	"fmt"
	"io"
	"os"

	"github.com/mrcook/svgtoolkit/geopattern/colour"
	"github.com/mrcook/svgtoolkit/geopattern/seed"
	"github.com/mrcook/svgtoolkit/svg"
)

const defaultBackgroundColour = "#933c3c"

// Pattern represents a pattern state. It should be populated with the required
// packages for correct usage. It requires a svg.Canvas upon which to write the
// SVG shapes, along with basic colour and styling configuration.
type Pattern struct {
	Width  int // Width of the pattern canvas
	Height int // Height of the pattern canvas

	Seed       *seed.Seed      // Seed is used for select one of the generators.
	BaseColour *colour.Colour  // Base colour for the pattern.
	Styles     *StylePreset    // A set of styles for the pattern.
	Generators []GeneratorType // A list of generators to choose from.
	Svg        *svg.Canvas     // SVG document to write to.
}

// New returns a Pattern instance where the settings, such as colour, generator
// list, etc., are given a set of sensible defaults. The seed can be any string,
// such as a name.
func New(width, height int, seedValue string, image io.Writer) (*Pattern, error) {
	if seedValue == "" {
		return nil, fmt.Errorf("seed string required")
	}

	// tessellation requires at least 100
	if width < 100 || height < 100 {
		return nil, fmt.Errorf("both width and height must be >= 100")
	}

	bgColour, _ := colour.NewWithSeed(defaultBackgroundColour, seedValue)

	pattern := &Pattern{
		Width:  width,
		Height: height,

		Seed:       seed.New(seedValue),
		BaseColour: bgColour,
		Styles:     DefaultStylePresets(),
		Generators: DefaultGenerators,
		Svg:        svg.New(width, height, image),
	}
	return pattern, nil
}

// Generate is the function which produces a coloured geo pattern.
//
// Before executing this will check that all dependencies are satisfied. If
// possible defaults will be used, although Seed and Svg will return an error.
func (p Pattern) Generate() error {
	if p.Seed == nil {
		return fmt.Errorf("seed must be initialized")
	}
	if p.Svg == nil {
		return fmt.Errorf("svg must be initialized")
	}
	if p.BaseColour == nil {
		p.BaseColour, _ = colour.New(defaultBackgroundColour)
	}
	if len(p.Generators) == 0 {
		p.Generators = DefaultGenerators
	}
	if p.Styles == nil {
		p.Styles = DefaultStylePresets()
	}

	// generate background plane
	p.Svg.Rect(0, 0, p.Width, p.Height, fmt.Sprintf(`fill="%s"`, p.Svg.RGB(p.BaseColour.Rgb())))

	if err := p.generatePatterns(); err != nil {
		return err
	}

	return nil
}

// Selects the pattern generator based on the seed and executes it.
func (p Pattern) generatePatterns() error {
	// trigger pattern generator: chevron, etc.
	generatorID := p.determineGeneratorIndex()
	switch generatorID {
	case Chevrons:
		p.Chevrons()
	case CirclesConcentric:
		p.CirclesConcentric()
	case CirclesOverlapping:
		p.CirclesOverlapping()
	case Diamonds:
		p.Diamonds()
	case Hexagons:
		p.Hexagons()
	case Octagons:
		p.Octagons()
	case Plaid:
		p.Plaid()
	case PlusSigns:
		p.PlusSigns()
	case RingsOverlapping:
		p.RingsOverlapping()
	case SineWaves:
		os.Exit(1)
		p.SineWaves()
	case Squares:
		p.Squares()
	case SquaresMosaic:
		p.SquaresMosaic()
	case SquaresNested:
		p.SquaresNested()
	case Tessellation:
		p.Tessellation()
	case Triangles:
		p.Triangles()
	case Xes:
		p.Xes()
	default:
		return fmt.Errorf("invalid generator ID: %d", generatorID)
	}
	return nil
}

func (p Pattern) reMap(value, vMin, vMax, dMin, dMax float64) float64 {
	return colour.ReMap(value, vMin, vMax, dMin, dMax)
}

func (p Pattern) seedToInt(index, len int) float64 {
	return float64(p.Seed.ToInt(index, len))
}

func (p Pattern) fillColour(val int) string {
	if val%2 == 0 {
		return p.Styles.FillColourLight
	} else {
		return p.Styles.FillColourDark
	}
}

func (p Pattern) opacity(colourValue float64) float64 {
	return colour.ReMap(colourValue, 0, 15, p.Styles.OpacityMin, p.Styles.OpacityMax)
}

func (p Pattern) determineGeneratorIndex() GeneratorType {
	index := p.Seed.ToInt(20, 1)
	if index >= len(p.Generators) {
		return Xes
	}

	return p.Generators[index]
}

// GeneratorType identifies the type of pattern generator.
type GeneratorType int

// The complete list of available generators.
const (
	Chevrons GeneratorType = iota
	CirclesConcentric
	CirclesOverlapping
	Diamonds
	Hexagons
	Octagons
	Plaid
	PlusSigns
	RingsOverlapping
	SineWaves
	Squares
	SquaresMosaic
	SquaresNested
	Tessellation
	Triangles
	Xes
)

// DefaultGenerators is the complete list of generators.
var DefaultGenerators = []GeneratorType{
	Chevrons, CirclesConcentric, Diamonds, Hexagons, SquaresMosaic, SquaresNested, Octagons, CirclesOverlapping,
	RingsOverlapping, Plaid, PlusSigns, SineWaves, Squares, Tessellation, Triangles, Xes,
}
