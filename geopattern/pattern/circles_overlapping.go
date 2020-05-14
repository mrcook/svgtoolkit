package pattern

import (
	"fmt"
	"math"
)

// CirclesOverlapping generator pattern.
func (p Pattern) CirclesOverlapping() {
	scale := p.seedToInt(0, 1)
	diameter := p.reMap(scale, 0, 15, 25, 200)
	radius := diameter / 2

	// calculate new circles size to fit canvas width
	columns := math.Ceil(float64(p.Width) / radius)
	// update radius with new sizes
	radius = float64(p.Width) / columns
	// set correct number of columns
	cols := int(columns)
	rows := int(float64(p.Height)/radius) + 1

	index := 0
	for y := 0; y <= rows; y++ {
		for x := 0; x <= cols; x++ {
			colourValue := p.seedToInt(index, 1)

			var styles []string
			styles = append(styles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
			styles = append(styles, p.Svg.Style(fmt.Sprintf(`opacity: %f`, p.opacity(colourValue))))

			p.Svg.Circle(float64(x)*radius, float64(y)*radius, radius, styles...)

			// Add an extra one at top-right, for tiling.
			if x == 0 {
				p.Svg.Circle(6*radius, float64(y)*radius, radius, styles...)
			}

			// Add an extra row at the end that matches the first row, for tiling.
			if y == 0 {
				p.Svg.Circle(float64(x)*radius, 6*radius, radius, styles...)
			}

			// Add an extra one at bottom-right, for tiling.
			if x == 0 && y == 0 {
				p.Svg.Circle(6*radius, 6*radius, radius, styles...)
			}

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}
