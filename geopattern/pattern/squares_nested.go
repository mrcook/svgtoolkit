package pattern

import (
	"fmt"
	"math"
)

// SquaresNested generator pattern.
func (p *Pattern) SquaresNested() {
	blockSize := p.reMap(p.seedToInt(0, 1), 0, 15, 4, 12)
	squareSize := blockSize * 7

	// calculate new sizes to fit canvas width
	shapeSize := blockSize*2 + squareSize
	columns := math.Ceil(float64(p.Width) / shapeSize)
	newSize := float64(p.Width) / columns
	blockSize = (newSize * 0.263172388) * 0.4223355 // new size as result from reMap.
	squareSize = blockSize * 7
	// set correct number of columns
	cols := int(columns)
	rows := int((float64(p.Height) / blockSize) * 0.15)

	// add an indent from the left and top
	offset := blockSize / 2

	index := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			colourValue := p.seedToInt(index, 1)

			p.Svg.Rect(
				int(float64(x)*squareSize+float64(x)*blockSize*2+blockSize/2+offset),
				int(float64(y)*squareSize+float64(y)*blockSize*2+blockSize/2+offset),
				int(squareSize),
				int(squareSize),
				[]string{
					fmt.Sprintf(`fill="none"`),
					fmt.Sprintf(`stroke="%s"`, p.fillColour(int(colourValue))),
					p.Svg.Style(fmt.Sprintf(`opacity: %f`, p.opacity(colourValue)), fmt.Sprintf(`stroke-width: %f`, blockSize)),
				}...,
			)

			colourValue = p.seedToInt(39-index, 1)

			p.Svg.Rect(
				int(float64(x)*squareSize+float64(x)*blockSize*2+blockSize/2+blockSize*2+offset),
				int(float64(y)*squareSize+float64(y)*blockSize*2+blockSize/2+blockSize*2+offset),
				int(blockSize*3),
				int(blockSize*3),
				[]string{
					fmt.Sprintf(`fill="none"`),
					fmt.Sprintf(`stroke="%s"`, p.fillColour(int(colourValue))),
					p.Svg.Style(fmt.Sprintf(`opacity: %f`, p.opacity(colourValue)), fmt.Sprintf(`stroke-width: %fpx`, blockSize)),
				}...,
			)

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}
