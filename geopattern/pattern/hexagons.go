package pattern

import (
	"fmt"
	"math"
)

// Hexagons generator pattern.
func (p Pattern) Hexagons() {
	scale := p.seedToInt(0, 1)
	sideLen := p.reMap(scale, 0, 15, 8, 60)
	height := sideLen * math.Sqrt(3)
	width := sideLen * 2

	// adjust diamond size to fit canvas width
	columns := math.Ceil(float64(p.Width) / width)
	width = float64(p.Width) / columns
	cols := int(columns * 1.3)
	rows := int(float64(p.Height)/height) + 1

	index := 0
	for y := 0; y <= rows; y++ {
		for x := 0; x <= cols; x++ {
			colourValue := p.seedToInt(index, 1)

			var defaultStyles []string
			defaultStyles = append(defaultStyles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`fill-opacity="%.3f"`, p.opacity(colourValue)))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))

			dy := float64(y)*height + height/2
			if x%2 == 0 {
				dy = float64(y) * height
			}

			styles := defaultStyles
			styles = append(styles, p.Svg.Transform(p.Svg.Translate(float64(x)*sideLen*1.5-width/2, dy-height/2)))
			p.buildHexagonShape(sideLen, styles)

			// Add an extra one at top-right, for tiling.
			if x == 0 {
				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(p.Svg.Translate(6*sideLen*1.5-width/2, dy-height/2)))
				p.buildHexagonShape(sideLen, styles)
			}

			// Add an extra row at the end that matches the first row, for tiling.
			if y == 0 {
				dy = 6*height + height/2
				if x%2 == 0 {
					dy = 6 * height
				}

				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(p.Svg.Translate(float64(x)*sideLen*1.5-width/2, dy-height/2)))
				p.buildHexagonShape(sideLen, styles)
			}

			// Add an extra one at bottom-right, for tiling.
			if x == 0 && y == 0 {
				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(p.Svg.Translate(6*sideLen*1.5-width/2, 5*height+height/2)))
				p.buildHexagonShape(sideLen, styles)
			}

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}

func (p Pattern) buildHexagonShape(sideLen float64, styles []string) {
	c := sideLen
	a := c / 2
	b := math.Sin(60*math.Pi/180) * c
	points := [][2]float64{{0, b}, {a, 0}, {a + c, 0}, {2 * c, b}, {a + c, 2 * b}, {a, 2 * b}, {0, b}}

	p.Svg.Polyline(points, styles...)
}
