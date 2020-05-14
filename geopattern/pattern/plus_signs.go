package pattern

import "fmt"

// PlusSigns generator pattern.
func (p Pattern) PlusSigns() {
	squareSize := p.reMap(p.seedToInt(0, 1), 0, 15, 10, 25)
	plusSize := squareSize * 3

	// create a tiling pattern
	p.Svg.Defs()
	p.Svg.Pattern("pattern", 0.0, 0.0, squareSize*12, squareSize*12, true)
	p.Svg.Group()
	p.buildPlusSignsPattern(squareSize, plusSize)
	p.Svg.GroupClose()
	p.Svg.PatternClose()
	p.Svg.DefsClose()

	p.Svg.Rect(0, 0, p.Width, p.Height, `fill="url(#pattern)"`)
}

func (p Pattern) buildPlusSignsPattern(squareSize, plusSize float64) {
	index := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {
			colourValue := p.seedToInt(index, 1)

			dx := 1.0
			if y%2 == 0 {
				dx = 0
			}

			var defaultStyles []string
			defaultStyles = append(defaultStyles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
			defaultStyles = append(defaultStyles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))
			defaultStyles = append(defaultStyles, p.Svg.Style(fmt.Sprintf(`fill-opacity: %.3f`, p.opacity(colourValue))))

			styles := defaultStyles
			styles = append(styles, p.Svg.Transform(p.Svg.Translate(float64(x)*plusSize-float64(x)*squareSize+dx*squareSize-squareSize, float64(y)*plusSize-float64(y)*squareSize-plusSize/2)))
			p.buildPlusSignShape(squareSize, styles...)

			// Add an extra column on the right for tiling.
			if x == 0 {
				group := defaultStyles
				group = append(group, p.Svg.Transform(p.Svg.Translate(4*plusSize-float64(x)*squareSize+dx*squareSize-squareSize, float64(y)*plusSize-float64(y)*squareSize-plusSize/2)))
				p.buildPlusSignShape(squareSize, group...)
			}

			// Add an extra row on the bottom that matches the first row, for tiling.
			if y == 0 {
				group := defaultStyles
				group = append(group, p.Svg.Transform(p.Svg.Translate(float64(x)*plusSize-float64(x)*squareSize+dx*squareSize-squareSize, 4*plusSize-float64(y)*squareSize-plusSize/2)))
				p.buildPlusSignShape(squareSize, group...)
			}

			// Add an extra one at top-right and bottom-right, for tiling.
			if x == 0 && y == 0 {
				group := defaultStyles
				group = append(group, p.Svg.Transform(p.Svg.Translate(4*plusSize-float64(x)*squareSize+dx*squareSize-squareSize, 4*plusSize-float64(y)*squareSize-plusSize/2)))
				p.buildPlusSignShape(squareSize, group...)
			}
			index += 1
		}
	}
}

func (p Pattern) buildPlusSignShape(size float64, styles ...string) {
	p.Svg.Group(styles...)
	p.Svg.Rect(size, 0.0, size, size*3)
	p.Svg.Rect(0.0, size, size*3, size)
	p.Svg.GroupClose()
}
