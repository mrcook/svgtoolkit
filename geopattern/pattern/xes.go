package pattern

import "fmt"

// Xes generator pattern.
func (p Pattern) Xes() {
	squareSize := p.reMap(p.seedToInt(0, 1), 0, 15, 10, 25)
	xSize := squareSize * 3 * 0.943

	// create a tiling pattern
	p.Svg.Defs()
	p.Svg.Pattern("pattern", 0.0, 0.0, xSize*3, xSize*3, true)
	p.Svg.Group()
	p.buildXesPattern(squareSize, xSize)
	p.Svg.GroupClose()
	p.Svg.PatternClose()
	p.Svg.DefsClose()

	p.Svg.Rect(0, 0, p.Width, p.Height, `fill="url(#pattern)"`)
}

func (p Pattern) buildXesPattern(squareSize, xSize float64) {
	index := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {
			colourValue := p.seedToInt(index, 1)

			var dy float64
			if x%2 == 0 {
				dy = float64(y)*xSize - xSize*0.5
			} else {
				dy = float64(y)*xSize - xSize*0.5 + xSize/4
			}

			var defaultStyles []string
			defaultStyles = append(defaultStyles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
			defaultStyles = append(defaultStyles, p.Svg.Style(fmt.Sprintf(`opacity: %.3f`, p.opacity(colourValue))))

			transforms := []string{
				p.Svg.Translate(float64(x)*xSize/2-xSize/2, dy-float64(y)*xSize/2),
				p.Svg.Rotate(45.0, xSize/2, xSize/2),
			}
			styles := defaultStyles
			styles = append(styles, p.Svg.Transform(transforms...))
			p.buildXesPlusShape(squareSize, styles)

			// Add an extra column on the right for tiling.
			if x == 0 {
				transforms = []string{
					p.Svg.Translate(6*xSize/2-xSize/2, dy-float64(y)*xSize/2),
					p.Svg.Rotate(45.0, xSize/2, xSize/2),
				}
				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(transforms...))
				p.buildXesPlusShape(squareSize, styles)
			}

			// Add an extra row on the bottom that matches the first row, for tiling.
			if y == 0 {
				if x%2 == 0 {
					dy = 6*xSize - xSize/2
				} else {
					dy = 6*xSize - xSize/2 + xSize/4
				}

				transforms = []string{
					p.Svg.Translate(float64(x)*xSize/2-xSize/2, dy-6*xSize/2),
					p.Svg.Rotate(45.0, xSize/2, xSize/2),
				}
				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(transforms...))
				p.buildXesPlusShape(squareSize, styles)
			}

			// These can hang off the bottom, so put a row at the top for tiling.
			if y == 5 {
				transforms = []string{
					p.Svg.Translate(float64(x)*xSize/2-xSize/2, dy-11*xSize/2),
					p.Svg.Rotate(45.0, xSize/2, xSize/2),
				}
				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(transforms...))
				p.buildXesPlusShape(squareSize, styles)
			}

			// Add an extra one at top-right and bottom-right, for tiling.
			if x == 0 && y == 0 {
				transforms = []string{
					p.Svg.Translate(6*xSize/2-xSize/2, dy-6*xSize/2),
					p.Svg.Rotate(45.0, xSize/2, xSize/2),
				}
				styles = defaultStyles
				styles = append(styles, p.Svg.Transform(transforms...))
				p.buildXesPlusShape(squareSize, styles)
			}
			index += 1
		}
	}
}

func (p Pattern) buildXesPlusShape(squareSize float64, styles []string) {
	p.Svg.Group(styles...)
	p.buildPlusSignShape(squareSize)
	p.Svg.GroupClose()
}
