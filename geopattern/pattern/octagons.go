package pattern

import (
	"fmt"
	"math"
)

// Octagons generator pattern.
func (p *Pattern) Octagons() {
	size := p.reMap(p.seedToInt(0, 1), 0, 15, 10, 60)

	// adjust diamond size to fit canvas width
	columns := math.Ceil(float64(p.Width) / size)
	size = float64(p.Width) / columns
	cols := int(columns)
	rows := int(float64(p.Height) / size)

	// pre-build Octagon shape
	s := size
	c := s * 0.33
	tile := [][2]float64{{c, 0}, {s - c, 0}, {s, c}, {s, s - c}, {s - c, s}, {c, s}, {0, s - c}, {0, c}, {c, 0}}

	index := 0
	for y := 0; y <= rows; y++ {
		for x := 0; x < cols; x++ {
			colourValue := p.seedToInt(index, 1)

			p.Svg.Polyline(tile,
				[]string{
					fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))),
					fmt.Sprintf(`fill-opacity="%f"`, p.opacity(colourValue)),
					fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour),
					fmt.Sprintf(`stroke-opacity="%f"`, p.Styles.StrokeOpacity),
					p.Svg.Transform(p.Svg.Translate(float64(x)*size, float64(y)*size)),
				}...,
			)

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}
