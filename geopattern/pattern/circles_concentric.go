package pattern

import (
	"fmt"
	"math"
)

// CirclesConcentric generator pattern.
func (p Pattern) CirclesConcentric() {
	scale := p.seedToInt(0, 1)
	bringSize := p.reMap(scale, 0, 15, 10, 60)
	strokeWidth := bringSize / 5

	//// calculate new circles size to fit canvas width
	columns := math.Ceil(float64(p.Width) / (bringSize + strokeWidth))
	size := float64(p.Width) / columns
	// update ring and stroke sizes
	ringSize := size * 0.83333 // new size as result from reMap.
	strokeWidth = ringSize / 5
	// set correct number of columns
	cols := int(columns)
	rows := int(float64(p.Height) / ringSize)

	index := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			colourValue := p.seedToInt(index, 1)
			opacity := p.opacity(colourValue)
			fill := p.fillColour(int(colourValue))

			cx := float64(x)*ringSize + float64(x)*strokeWidth + (ringSize+strokeWidth)/2
			cy := float64(y)*ringSize + float64(y)*strokeWidth + (ringSize+strokeWidth)/2

			p.Svg.Circle(
				cx, cy, ringSize/2,
				[]string{
					fmt.Sprintf(`fill="none"`),
					fmt.Sprintf(`stroke="%s"`, fill),
					p.Svg.Style(fmt.Sprintf(`opacity: %f`, opacity), fmt.Sprintf(`stroke-width: %f`, strokeWidth)),
				}...,
			)

			colourValue = p.seedToInt(39-index, 1)
			opacity = p.opacity(colourValue)
			fill = p.fillColour(int(colourValue))

			p.Svg.Circle(
				cx, cy, ringSize/4,
				[]string{
					fmt.Sprintf(`fill="%s"`, fill),
					fmt.Sprintf(`fill-opacity="%f"`, opacity),
				}...,
			)

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}
