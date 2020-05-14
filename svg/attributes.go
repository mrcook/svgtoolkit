package svg

import (
	"fmt"
	"strings"
)

// Style allows attributes that correspond to those in HTML to be used in
// various SVG elements. One or more attributes can be given, which will be
// joined together using a semi-colon.
//
// SVG Reference: https://www.w3.org/TR/SVG11/styling.html#StyleElement
func (s Canvas) Style(attributes ...string) string {
	return fmt.Sprintf(`style="%s"`, strings.Join(attributes, ";"))
}

// Transform wraps a transform list and returns then as a string.
//
// The value of the transform attribute is a <transform-list>, which is defined
// as a list of transform definitions, which are applied in the order provided.
// Available transform definitions are: matrix, translate, scale, rotate, skewX, and skewY.
//
// SVG Reference: https://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (s Canvas) Transform(definitions ...string) string {
	return fmt.Sprintf(`transform="%s"`, strings.Join(definitions, " "))
}

// Translate definition which specifies a translation by tx and ty for use
// within the transform attribute.
func (s Canvas) Translate(x, y interface{}) string {
	return fmt.Sprintf("translate(%v, %v)", x, y)
}

// Scale definition which specifies a scale operation by sx and sy, for use
// within the transform attribute.
func (s Canvas) Scale(x, y interface{}) string {
	return fmt.Sprintf("scale(%v, %v)", x, y)
}

// Rotate definition which specifies a rotation by rotate-angle degrees about
// a given point, for use within the transform attribute.
func (s Canvas) Rotate(a, x, y interface{}) string {
	return fmt.Sprintf("rotate(%v, %v, %v)", a, x, y)
}

// ViewBox attribute can be used to allow a given set of graphics to be
// stretched to fit a particular container element.
//
// SVG Reference: https://www.w3.org/TR/SVG11/coords.html#ViewBoxAttribute
func (s Canvas) ViewBox(x, y, w, h int) string {
	return fmt.Sprintf(`viewBox="%d %d %d %d"`, x, y, w, h)
}

// RGB wraps the given values in a `rgb()` function.
func (s Canvas) RGB(r, g, b interface{}) string {
	return fmt.Sprintf(`rgb(%v, %v, %v)`, r, g, b)
}
