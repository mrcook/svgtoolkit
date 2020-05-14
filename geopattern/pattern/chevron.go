package pattern

import (
	"fmt"
	"math"
)

// Chevrons generator pattern.
func (p Pattern) Chevrons() {
	width := p.reMap(p.seedToInt(0, 1), 0, 15, 30, 80)
	height := p.reMap(p.seedToInt(0, 1), 0, 15, 30, 80)

	// adjust chevron width to fit canvas width
	columns := math.Ceil(float64(p.Width) / width)
	width = float64(p.Width) / columns

	cols := int(columns)
	rows := int(float64(p.Height) / height * 2)

	index := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			colourValue := p.seedToInt(index, 1)

			var styles []string
			styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
			styles = append(styles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))
			styles = append(styles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
			styles = append(styles, fmt.Sprintf(`fill-opacity="%.3f"`, p.opacity(colourValue)))
			styles = append(styles, `stroke-width="1"`)

			group := styles
			group = append(group, p.Svg.Transform(p.Svg.Translate(float64(x)*width, float64(y)*height*0.66-height/2)))
			p.Svg.Group(group...)
			p.buildChevronShape(width, height)
			p.Svg.GroupClose()

			// Add an extra row at the end that matches the first row, for tiling.
			if y == 0 {
				group = styles
				group = append(group, p.Svg.Transform(p.Svg.Translate(float64(x)*width, 6*height*0.66-height/2)))
				p.Svg.Group(group...)
				p.buildChevronShape(width, height)
				p.Svg.GroupClose()
			}
			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}

func (p Pattern) buildChevronShape(width, height float64) {
	e := height * 0.66

	p.Svg.Polyline([][2]float64{{0, 0}, {width / 2, height - e}, {width / 2, height}, {0, e}, {0, 0}})
	p.Svg.Polyline([][2]float64{{width / 2, height - e}, {width, 0}, {width, e}, {width / 2, height}, {width / 2, height - e}})
}
