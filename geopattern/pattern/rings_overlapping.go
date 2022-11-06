package pattern

import "fmt"

// RingsOverlapping generator pattern.
func (p *Pattern) RingsOverlapping() {
	scale := p.seedToInt(0, 1)
	size := p.reMap(scale, 0, 15, 10, 60)
	strokeWidth := size / 4

	// create a tiling pattern
	p.Svg.Defs()
	p.Svg.Pattern("pattern", 0.0, 0.0, size*6, size*6, true)
	p.Svg.Group()
	p.buildRingsOverlappingPattern(size, strokeWidth)
	p.Svg.GroupClose()
	p.Svg.PatternClose()
	p.Svg.DefsClose()

	p.Svg.Rect(0, 0, p.Width, p.Height, `fill="url(#pattern)"`)
}

func (p *Pattern) buildRingsOverlappingPattern(size, strokeWidth float64) {
	index := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {
			colourValue := p.seedToInt(index, 1)

			var styles []string
			styles = append(styles, fmt.Sprintf(`fill="none"`))
			styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.fillColour(int(colourValue))))
			styles = append(styles, p.Svg.Style(fmt.Sprintf(`opacity: %f`, p.opacity(colourValue)), fmt.Sprintf(`stroke-width: %fpx`, strokeWidth)))

			p.Svg.Circle(float64(x)*size, float64(y)*size, size-strokeWidth/2, styles...)

			// Add an extra one at top-right, for tiling.
			if x == 0 {
				p.Svg.Circle(6*size, float64(y)*size, size-strokeWidth/2, styles...)
			}

			if y == 0 {
				p.Svg.Circle(float64(x)*size, 6*size, size-strokeWidth/2, styles...)
			}

			if x == 0 && y == 0 {
				p.Svg.Circle(6*size, 6*size, size-strokeWidth/2, styles...)
			}
			index += 1
		}
	}
}
