package pattern

import (
	"fmt"
	"math"
)

// Diamonds generator pattern.
func (p *Pattern) Diamonds() {
	width := p.reMap(p.seedToInt(0, 1), 0, 15, 10, 50)
	height := p.reMap(p.seedToInt(1, 1), 0, 15, 10, 50)

	// adjust diamond size to fit canvas width
	columns := math.Ceil(float64(p.Width) / width)
	width = float64(p.Width) / columns
	cols := int(columns)
	rows := int(float64(p.Height)/height) * 2

	diamondPoints := [][2]float64{{width / 2, 0}, {width, height / 2}, {width / 2, height}, {0, height / 2}}

	index := 0
	for y := 0; y <= rows; y++ {
		for x := 0; x <= cols; x++ {
			colourValue := p.seedToInt(index, 1)

			dx := width / 2
			if y%2 == 0 {
				dx = 0
			}

			var styles []string
			styles = append(styles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
			styles = append(styles, fmt.Sprintf(`fill-opacity="%.3f"`, p.opacity(colourValue)))
			styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
			styles = append(styles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))

			group := styles
			group = append(group, p.Svg.Transform(p.Svg.Translate(float64(x)*width-width/2+dx, height/2*float64(y)-height/2)))
			p.Svg.Polyline(diamondPoints, group...)

			// Add an extra one at top-right, for tiling.
			if x == 0 {
				group = styles
				group = append(group, p.Svg.Transform(p.Svg.Translate(6*width-width/2+dx, height/2*float64(y)-height/2)))
				p.Svg.Polyline(diamondPoints, group...)
			}

			// Add an extra row at the end that matches the first row, for tiling.
			if y == 0 {
				group = styles
				group = append(group, p.Svg.Transform(p.Svg.Translate(float64(x)*width-width/2+dx, height/2*6-height/2)))
				p.Svg.Polyline(diamondPoints, group...)
			}

			// Add an extra one at bottom-right, for tiling.
			if x == 0 && y == 0 {
				group = styles
				group = append(group, p.Svg.Transform(p.Svg.Translate(6*width-width/2+dx, height/2*6-height/2)))
				p.Svg.Polyline(diamondPoints, group...)
			}

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}
