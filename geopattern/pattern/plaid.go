package pattern

import "fmt"

// Plaid generator pattern.
func (p *Pattern) Plaid() {
	height := 0.0
	width := 0.0

	// fill the width and height with shapes
	rows := p.Height / 25
	cols := p.Width / 25

	// horizontal stripes
	index := 0
	for l := 0; l < rows; l++ {
		space := p.seedToInt(index, 1)
		height += space + 5

		colourValue := p.seedToInt(index+1, 1)
		stripeHeight := colourValue + 5

		p.Svg.Rect(
			0.0, height,
			float64(p.Width), stripeHeight,
			[]string{
				fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))),
				fmt.Sprintf(`opacity="%f"`, p.opacity(colourValue)),
			}...,
		)
		height += stripeHeight
		index += 2
		if index > 38 {
			index = 0
		}
	}

	// vertical stripes
	index = 0
	for l := 0; l < cols; l++ {
		space := p.seedToInt(index, 1)
		width += space + 5

		colourValue := p.seedToInt(index+1, 1)
		stripeWidth := colourValue + 5

		p.Svg.Rect(
			width, 0.0,
			stripeWidth, float64(p.Height),
			[]string{
				fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))),
				fmt.Sprintf(`opacity="%f"`, p.opacity(colourValue)),
			}...,
		)
		width += stripeWidth
		index += 2
		if index > 38 {
			index = 0
		}
	}
}
