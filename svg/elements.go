package svg

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Circle defines a circle based on a center point and a radius.
//
// SVG Reference: https://www.w3.org/TR/SVG11/shapes.html#CircleElement
func (s *Canvas) Circle(cx, cy, r interface{}, attributes ...string) {
	s.printf(`<circle cx="%v" cy="%v" r="%v" %s />`, cx, cy, r, strings.Join(attributes, " "))
	s.println()
}

// Defs element is a container element for referenced elements.
//
// SVG Reference: http://www.w3.org/TR/SVG11/struct.html#DefsElement
func (s *Canvas) Defs()      { s.println(`<defs>`) }
func (s *Canvas) DefsClose() { s.println(`</defs>`) }

// Filter
// SVG Reference: http://www.w3.org/TR/SVG11/filters.html#FilterElement
func (s *Canvas) Filter(id string, attributes ...string) {
	s.printf(`<filter id="%s" %s>`, id, strings.Join(attributes, " "))
	s.println()
}

// FeBlend composites two objects together using commonly used imaging software blending modes.
// It performs a pixel-wise combination of two input images.
//
// in = first input image to the blending operation.
// in2 = second input image to the blending operation.
// mode = the image blending mode: normal, multiply, screen, darken, lighten.
//
// SVG Reference: https://www.w3.org/TR/SVG11/filters.html#feBlendElement
func (s *Canvas) FeBlend(in, in2, mode string, attributes ...string) {
	mode = strings.ToLower(mode)

	switch mode {
	case "normal", "multiply", "screen", "darken", "lighten":
		// valid mode type
	default:
		mode = "normal"
	}

	s.printf(`<feBlend in="%s" in2="%s" mode="%s" %s />`, in, in2, mode, strings.Join(attributes, " "))
	s.println()
}

// FeGaussianBlur provides a filter primitive to perform a Gaussian blur on the input image.
//
// SVG Reference: http://www.w3.org/TR/SVG11/filters.html#feGaussianBlurElement
func (s *Canvas) FeGaussianBlur(in string, stdDeviationX, stdDeviationY int, attributes ...string) {
	// negative numbers are an error, so prevent it
	if stdDeviationX < 0 {
		stdDeviationX = 0
	}
	if stdDeviationY < 0 {
		stdDeviationY = 0
	}
	s.printf(`<feGaussianBlur in="%s" stdDeviation="%d %d" %s />`, in, stdDeviationX, stdDeviationY, strings.Join(attributes, " "))
	s.println()
}

// FeOffset filter primitive offsets the input image relative to its current
// position in the image space by the specified vector.
//
// SVG Reference: http://www.w3.org/TR/SVG11/filters.html#feOffsetElement
func (s *Canvas) FeOffset(in string, dx, dy interface{}, attributes ...string) {
	s.printf(`<feOffset in="%s" dx="%v" dy="%v" %s />`, in, dx, dy, strings.Join(attributes, " "))
	s.println()
}

// FilterClose writes the closing `</filter> tag to the canvas.
func (s *Canvas) FilterClose() {
	s.println(`</filter>`)
}

// Group element is a container element for grouping together related graphics elements.
//
// SVG Reference: https://www.w3.org/TR/SVG11/struct.html#Groups
func (s *Canvas) Group(attributes ...string) {
	s.printf(`<g %s>`, strings.Join(attributes, " "))
	s.println()
}

// GroupClose writes the closing `</g> tag to the canvas.
func (s *Canvas) GroupClose() { s.println(`</g>`) }

// Path represents the outline of a shape which can be filled, stroked, used
// as a clipping path, or any combination of the three.
//
// SVG Reference: https://www.w3.org/TR/SVG11/paths.html#PathElement
func (s *Canvas) Path(pathData string, attributes ...string) {
	s.printf(`<path d="%s" %s />`, pathData, strings.Join(attributes, " "))
	s.println()
}

// Pattern is used to fill or stroke an object using a pre-defined graphic
// object which can be replicated ("tiled") at fixed intervals in x and y
// to cover the areas to be painted.
//
// SVG Reference: http://www.w3.org/TR/SVG11/pservers.html#Patterns
func (s *Canvas) Pattern(id string, x, y, width, height interface{}, userSpace bool, attributes ...string) {
	patternUnits := "userSpaceOnUse"
	if !userSpace {
		patternUnits = "objectBoundingBox"
	}
	s.printf(`<pattern id="%s" x="%v" y="%v" width="%v" height="%v" patternUnits="%s" %s>`, id, x, y, width, height, patternUnits, strings.Join(attributes, " "))
	s.println()
}

// PatternClose writes the closing `</pattern> tag to the canvas.
func (s *Canvas) PatternClose() { s.println("</pattern>") }

// Polygon element defines a closed shape consisting of a set of connected straight line segments.
//
// SVG Reference: http://www.w3.org/TR/SVG11/shapes.html#PolygonElement
func (s *Canvas) Polygon(coords [][2]float64, attributes ...string) {
	var points []string

	for _, p := range coords {
		points = append(points, fmt.Sprintf("%f,%f", p[0], p[1]))
	}

	s.printf(`<polygon points="%s" %s />`, strings.Join(points, " "), strings.Join(attributes, " "))
}

// Polyline element defines a set of connected straight line segments.
// Typically, polyline elements define open shapes.
//
// SVG Reference: https://www.w3.org/TR/SVG11/shapes.html#PolylineElement
func (s *Canvas) Polyline(coords [][2]float64, attributes ...string) {
	var points []string

	for _, p := range coords {
		points = append(points, fmt.Sprintf("%f,%f", p[0], p[1]))
	}

	s.printf(`<polyline points="%s" %s />`, strings.Join(points, " "), strings.Join(attributes, " "))
	s.println()
}

// Rect element defines a rectangle, placed at the requested coordinates using
// the dimensions.
//
// SVG Reference: http://www.w3.org/TR/SVG11/shapes.html#RectElement
func (s *Canvas) Rect(x, y, w, h interface{}, attributes ...string) {
	s.printf(`<rect x="%v" y="%v" width="%v" height="%v" %s />`, x, y, w, h, strings.Join(attributes, " "))
	s.println()
}

// RectRounded element defines a rounded rectangle, placed at the requested
// coordinates using the dimensions. Rounded-ness can be achieved by setting
// appropriate values for attributes ‘rx’ and ‘ry’.
//
// SVG Reference: http://www.w3.org/TR/SVG11/shapes.html#RectElement
func (s *Canvas) RectRounded(x, y, w, h, rx, ry interface{}, attributes ...string) {
	s.printf(`<rect x="%v" y="%v" width="%v" height="%v" rx="%v" ry="%v"  %s />`, x, y, w, h, rx, ry, strings.Join(attributes, " "))
	s.println()
}

// Text element defines a graphics element consisting of text, placed at the requested coordinates.
//
// SVG Reference: http://www.w3.org/TR/SVG11/text.html#TextElement
func (s *Canvas) Text(x int, y int, text string, attributes ...string) {
	s.printf(`<text x="%d" y="%d" %s>`, x, y, strings.Join(attributes, " "))
	xml.Escape(s.writer, []byte(text))
	s.println("</text>")
}

// Use references another element and indicates that the graphical contents of
// that element is included/drawn at that given point in the document.
//
// SVG Reference: https://www.w3.org/TR/SVG11/struct.html#UseElement
func (s *Canvas) Use(xlink string, attributes ...string) {
	s.printf(`<use xlink:href="#%s" %s />`, xlink, strings.Join(attributes, " "))
	s.println()
}

//
// Helper functions for working with points in polygon/polyline elements.
//

func (s *Canvas) mapCoordsList(coords [][2]interface{}) []string {
	var points []string

	for i := 0; i < len(coords); i++ {
		pointX, xOk := s.interfaceNumberToString(coords[i][0])
		pointY, yOk := s.interfaceNumberToString(coords[i][1])

		if !xOk || !yOk {
			return []string{"-- ERROR: only int/float64 point types allowed --"}
		}

		points = append(points, fmt.Sprintf("%s,%s", pointX, pointY))
	}

	return points
}

func (s *Canvas) interfaceNumberToString(num interface{}) (string, bool) {
	switch num.(type) {
	case int:
		return fmt.Sprintf("%f", num), true
	case float64:
		return fmt.Sprintf("%f", num), true
	default:
		return "", false
	}
}
