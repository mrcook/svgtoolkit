package pattern

import (
	"fmt"
	"math"
)

// Triangles generator pattern.
func (p Pattern) Triangles() {
	scale := p.seedToInt(0, 1)
	sideLen := p.reMap(scale, 0, 15, 15, 80)
	height := sideLen / 2 * math.Sqrt(3)

	// calculate new sizes to fit canvas width
	columns := math.Ceil(float64(p.Width) / sideLen)
	sideLen = float64(p.Width) / columns
	height = sideLen / 2 * math.Sqrt(3)
	// set correct number of columns
	cols := int(columns * 2)
	rows := int(float64(p.Height) / height)

	index := 0
	for y := 0; y <= rows; y++ {
		for x := 0; x <= cols; x++ {
			colourValue := p.seedToInt(index, 1)

			var defaultStyles []string
			defaultStyles = append(defaultStyles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`fill-opacity="%.3f"`, p.opacity(colourValue)))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))

			rotation := 0.0
			if y%2 == 0 {
				if x%2 == 0 {
					rotation = 180
				}
			} else {
				if x%2 != 0 {
					rotation = 180
				}
			}

			transforms := []string{
				p.Svg.Translate(float64(x)*sideLen*0.5-sideLen/2, height*float64(y)),
				p.Svg.Rotate(rotation, sideLen/2, height/2),
			}
			styles := defaultStyles
			styles = append(styles, p.Svg.Transform(transforms...))
			p.buildTriangleShape(sideLen, height, styles)

			// Add an extra one at top-right, for tiling.
			if x == 0 {
				transforms = []string{
					p.Svg.Translate(6*sideLen*0.5-sideLen/2, height*float64(y)),
					p.Svg.Rotate(rotation, sideLen/2, height/2),
				}
				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(transforms...))
				p.buildTriangleShape(sideLen, height, styles)
			}

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}

func (p Pattern) buildTriangleShape(sideLen, height float64, styles []string) {
	halfWidth := sideLen / 2
	points := [][2]float64{{halfWidth, 0}, {sideLen, height}, {0, height}, {halfWidth, 0}}

	p.Svg.Polyline(points, styles...)
}
