package pattern

import (
	"fmt"
	"math"
)

// Squares generator pattern.
func (p *Pattern) Squares() {
	size := p.reMap(p.seedToInt(0, 1), 0, 15, 10, 60)

	// fill the canvas with shapes
	columns := math.Ceil(float64(p.Width) / size)
	size = float64(p.Width) / columns

	cols := int(columns)
	rows := int(float64(p.Height) / size)

	index := 0
	for y := 0; y <= rows; y++ {
		for x := 0; x < cols; x++ {
			colourValue := p.seedToInt(index, 1)

			p.Svg.Rect(
				int(float64(x)*size), int(float64(y)*size),
				int(size), int(size),
				[]string{
					fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))),
					fmt.Sprintf(`fill-opacity="%f"`, p.opacity(colourValue)),
					fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour),
					fmt.Sprintf(`stroke-opacity="%f"`, p.Styles.StrokeOpacity),
				}...,
			)

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}
