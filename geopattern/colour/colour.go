// Package colour is used for generating colours for the patterns.
package colour

import (
	"github.com/lucasb-eyer/go-colorful"

	"github.com/mrcook/svgtoolkit/geopattern/seed"
)

// Colour represents a background colour for the generated pattern.
type Colour struct {
	seed   *seed.Seed
	colour colorful.Color
}

// New returns a Colour instance. It takes a HTML colour as a hex string: `#ff00ff`.
func New(htmlColour string) (*Colour, error) {
	c, err := colorful.Hex(htmlColour)
	if err != nil {
		return nil, err
	}

	return &Colour{colour: c}, nil
}

// New returns a Colour instance. It takes a HTML colour as a hex string: `#ff00ff`,
// and string to be used as a seed for generating a different colour palette.
func NewWithSeed(htmlColour string, str string) (*Colour, error) {
	c, err := colorful.Hex(htmlColour)
	if err != nil {
		return nil, err
	}

	colour := &Colour{
		seed:   seed.New(str),
		colour: c,
	}
	return colour, nil
}

// HtmlHex returns a HTML hex colour string: `#00ff00`.
func (c *Colour) HtmlHex() string {
	if c.seed != nil {
		c.transform()
	}
	return c.colour.Hex()
}

// Rgb returns the RGB integer values for each colour.
func (c Colour) Rgb() (int, int, int) {
	if c.seed != nil {
		c.transform()
	}

	r := int(c.colour.R * 255)
	g := int(c.colour.G * 255)
	b := int(c.colour.B * 255)

	return r, g, b
}

// Transform adjusts the colours' hue and saturation values using the seed.
func (c *Colour) transform() {
	hue, sat, lum := c.colour.Hsl()

	hueSeedOffset := float64(c.seed.ToInt(14, 3))
	hueOffset := ReMap(hueSeedOffset, 0, 4095, 0, 359)
	hue = hue - hueOffset

	satOffset := float64(c.seed.ToInt(17, 1))

	// sat as a percentage is 1.0, not 100, so fix offset value
	if int(satOffset)%2 == 0 {
		sat += satOffset / 100
	} else {
		sat -= satOffset / 100
	}
	c.colour = colorful.Hsl(hue, sat, lum)
}

// ReMap is a Go implementation of Processing's map function,
// which maps a number from one range to another.
// Parameters:
//   - value: value to be converted.
//   - vMin/vMax: lower/upper bound of the value's current range.
//   - dMin/dMax: lower/upper bound of the value's target range.
//
// Reference: http://processing.org/reference/map_.html
func ReMap(value, vMin, vMax, dMin, dMax float64) float64 {
	vRange := vMax - vMin
	dRange := dMax - dMin
	return (value-vMin)*dRange/vRange + dMin
}
