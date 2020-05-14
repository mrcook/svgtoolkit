package pattern

import (
	"fmt"
	"math"
)

// SquaresMosaic generator pattern.
func (p Pattern) SquaresMosaic() {
	size := p.reMap(p.seedToInt(0, 1), 0, 15, 15, 50)

	// fill the canvas with shapes
	columns := math.Ceil(float64(p.Width) / size)
	size = float64(p.Width) / columns
	cols := int(columns)
	rows := int(float64(p.Height) / size)

	index := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if x%2 == 0 {
				if y%2 == 0 {
					p.drawOuterMosaicTile(float64(x)*size*2, float64(y)*size*2, size, p.seedToInt(index, 1))
				} else {
					p.drawInnerMosaicTile(float64(x)*size*2, float64(y)*size*2, size, [2]float64{p.seedToInt(index, 1), p.seedToInt(index+1, 1)})
				}
			} else {
				if y%2 == 0 {
					p.drawInnerMosaicTile(float64(x)*size*2, float64(y)*size*2, size, [2]float64{p.seedToInt(index, 1), p.seedToInt(index+1, 1)})
				} else {
					p.drawOuterMosaicTile(float64(x)*size*2, float64(y)*size*2, size, p.seedToInt(index, 1))
				}
			}

			index += 1
			if index > 38 {
				index = 0
			}
		}
	}
}

func (p Pattern) drawInnerMosaicTile(x, y, size float64, colourValues [2]float64) {
	triangle := buildRightTriangleShapePoints(size)

	var styles []string
	styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
	styles = append(styles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))
	styles = append(styles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValues[0]))))
	styles = append(styles, fmt.Sprintf(`fill-opacity="%.3f"`, p.opacity(colourValues[0])))

	group := styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x+size, y), p.Svg.Scale(-1, 1)))
	p.Svg.Polyline(triangle, group...)

	group = styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x+size, y+size*2), p.Svg.Scale(1, -1)))
	p.Svg.Polyline(triangle, group...)

	styles = []string{}
	styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
	styles = append(styles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))
	styles = append(styles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValues[1]))))
	styles = append(styles, fmt.Sprintf(`fill-opacity="%.3f"`, p.opacity(colourValues[1])))

	group = styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x+size, y+size*2), p.Svg.Scale(-1, -1)))
	p.Svg.Polyline(triangle, group...)

	group = styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x+size, y), p.Svg.Scale(1, 1)))
	p.Svg.Polyline(triangle, group...)
}

func (p Pattern) drawOuterMosaicTile(x, y, size, colourValue float64) {
	triangle := buildRightTriangleShapePoints(size)

	var styles []string
	styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
	styles = append(styles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))
	styles = append(styles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
	styles = append(styles, fmt.Sprintf(`fill-opacity="%.3f"`, p.opacity(colourValue)))

	group := styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x, y+size), p.Svg.Scale(1, -1)))
	p.Svg.Polyline(triangle, group...)

	group = styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x+size*2, y+size), p.Svg.Scale(-1, -1)))
	p.Svg.Polyline(triangle, group...)

	group = styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x, y+size), p.Svg.Scale(1, 1)))
	p.Svg.Polyline(triangle, group...)

	group = styles
	group = append(group, p.Svg.Transform(p.Svg.Translate(x+size*2, y+size), p.Svg.Scale(-1, 1)))
	p.Svg.Polyline(triangle, group...)
}

func buildRightTriangleShapePoints(sideLen float64) [][2]float64 {
	return [][2]float64{{0, 0}, {sideLen, sideLen}, {0, sideLen}, {0, 0}}
}
