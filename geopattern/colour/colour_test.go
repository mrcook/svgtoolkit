package colour_test

import (
	"testing"

	"github.com/mrcook/svgtoolkit/geopattern/colour"
)

func TestColour_RGB(t *testing.T) {
	c, _ := colour.New("#ff00ff")

	r, g, b := c.Rgb()
	if r != 255 || g != 0 || b != 255 {
		t.Fatalf("expected HTML colour to be parsed correctly, got: '%d, %d, %d'", r, g, b)
	}
}

func TestColour_ToHex(t *testing.T) {
	c, _ := colour.New("#00ff00")

	if c.HtmlHex() != ("#00ff00") {
		t.Fatalf("expected correct HTML hex to be returned, got: '%s'", c.HtmlHex())
	}
}

func TestColour_SeedAdjust(t *testing.T) {
	c, _ := colour.NewWithSeed("#ff00ff", "Seed Test")

	if c.HtmlHex() != ("#0be3f4") {
		t.Errorf("expected correct HTML colour to be returned, got: '%s'", c.HtmlHex())
	}

	r, g, b := c.Rgb()
	if r != 202 || g != 232 || b != 22 {
		t.Fatalf("expected correct RGB colour to be returned, got: '%d, %d, %d'", r, g, b)
	}
}

func TestMap(t *testing.T) {
	mapped := colour.ReMap(2, 0, 400, 0, 100)
	if mapped != 0.5 {
		t.Fatalf("expected values to be mapped correctly, got: %f", mapped)
	}

	mapped = colour.ReMap(2, 0, 15, 30, 80)
	if mapped != 36.666666666666664 {
		t.Fatalf("expected values to be mapped correctly, got: %v", mapped)
	}
}
