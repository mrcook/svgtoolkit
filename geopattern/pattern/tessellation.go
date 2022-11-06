package pattern

import (
	"fmt"
	"math"
)

// Tessellation generator pattern.
func (p *Pattern) Tessellation() {
	// 3.4.6.4 semi-regular tessellation
	sideLen := p.reMap(p.seedToInt(0, 1), 0, 15, 10, 35) // was dMin: 5, dMax: 40
	hexWidth := sideLen * 2
	hexHeight := sideLen * math.Sqrt(3)
	triangleHeight := sideLen / 2 * math.Sqrt(3)
	tileWidth := sideLen*3 + triangleHeight*2
	tileHeight := (hexHeight * 2) + (sideLen * 2)

	// create a tiling pattern
	p.Svg.Defs()
	p.Svg.Pattern("pattern", 0.0, 0.0, tileWidth, tileHeight, true)
	p.Svg.Group()
	p.buildTessellationPattern(sideLen, hexWidth, hexHeight, tileWidth, tileHeight, triangleHeight)
	p.Svg.GroupClose()
	p.Svg.PatternClose()
	p.Svg.DefsClose()

	p.Svg.Rect(0, 0, p.Width, p.Height, `fill="url(#pattern)"`)
}

func (p *Pattern) buildTessellationPattern(sideLen, hexWidth, hexHeight, tileWidth, tileHeight, triangleHeight float64) {
	// build rotated triangle shape points
	triangle := [][2]float64{{0, 0}, {triangleHeight, sideLen / 2}, {0, sideLen}, {0, 0}}

	var transforms []string

	for index := 0; index <= 19; index++ {
		val := p.seedToInt(index, 1)

		styles := p.tessellationStyles(val, "")

		switch index {
		case 0: // all 4 corners
			p.Svg.Rect(-sideLen/2, -sideLen/2, sideLen, sideLen, styles...)
			p.Svg.Rect(tileWidth-sideLen/2, -sideLen/2, sideLen, sideLen, styles...)
			p.Svg.Rect(-sideLen/2, tileHeight-sideLen/2, sideLen, sideLen, styles...)
			p.Svg.Rect(tileWidth-sideLen/2, tileHeight-sideLen/2, sideLen, sideLen, styles...)
		case 1: // center / top square
			p.Svg.Rect(hexWidth/2+triangleHeight, hexHeight/2, sideLen, sideLen, styles...)
		case 2: // side squares
			p.Svg.Rect(-sideLen/2, tileHeight/2-sideLen/2, sideLen, sideLen, styles...)
			p.Svg.Rect(tileWidth-sideLen/2, tileHeight/2-sideLen/2, sideLen, sideLen, styles...)
		case 3: // center / bottom square
			p.Svg.Rect(hexWidth/2+triangleHeight, hexHeight*1.5+sideLen, sideLen, sideLen, styles...)
		case 4: // left top / bottom triangle
			transforms = []string{
				p.Svg.Translate(sideLen/2, -sideLen/2),
				p.Svg.Rotate(0.0, sideLen/2, triangleHeight/2),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)

			transforms = []string{
				p.Svg.Translate(sideLen/2, tileHeight-sideLen/2),
				p.Svg.Rotate(0.0, sideLen/2, triangleHeight/2),
				p.Svg.Scale(1, -1),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 5: // right top / bottom triangle
			transforms = []string{
				p.Svg.Translate(tileWidth-sideLen/2, -sideLen/2),
				p.Svg.Rotate(0.0, sideLen/2, triangleHeight/2),
				p.Svg.Scale(-1, 1),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)

			transforms = []string{
				p.Svg.Translate(tileWidth-sideLen/2, tileHeight+sideLen/2),
				p.Svg.Rotate(0.0, sideLen/2, triangleHeight/2),
				p.Svg.Scale(-1, -1),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 6: // center / top / right triangle
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(p.Svg.Translate(tileWidth/2+sideLen/2, hexHeight/2)))...)
		case 7: // center / top / left triangle
			transforms = []string{
				p.Svg.Translate(tileWidth-tileWidth/2-sideLen/2, hexHeight/2),
				p.Svg.Scale(-1, 1),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 8: // center / bottom / right triangle
			transforms = []string{
				p.Svg.Translate(tileWidth/2+sideLen/2, tileHeight-hexHeight/2),
				p.Svg.Scale(1, -1),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 9: // center / bottom / left triangle
			transforms = []string{
				p.Svg.Translate(tileWidth-tileWidth/2-sideLen/2, tileHeight-hexHeight/2),
				p.Svg.Scale(-1, -1),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 10: // left / middle triangle
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(p.Svg.Translate(sideLen/2, tileHeight/2-sideLen/2)))...)
		case 11: // right / middle triangle
			transforms = []string{
				p.Svg.Translate(tileWidth-sideLen/2, tileHeight/2-sideLen/2),
				p.Svg.Scale(-1, 1),
			}
			p.Svg.Polyline(triangle, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 12: // left / top square
			transforms = []string{
				p.Svg.Translate(sideLen/2, sideLen/2),
				p.Svg.Rotate(-30.0, 0.0, 0.0),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 13: // right / top square
			transforms = []string{
				p.Svg.Scale(-1, 1),
				p.Svg.Translate(-tileWidth+sideLen/2, sideLen/2),
				p.Svg.Rotate(-30.0, 0.0, 0.0),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 14: // left / center-top square
			transforms = []string{
				p.Svg.Translate(sideLen/2, tileHeight/2-sideLen/2-sideLen),
				p.Svg.Rotate(30.0, 0.0, sideLen),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 15: // right / center-top square
			transforms = []string{
				p.Svg.Scale(-1, 1),
				p.Svg.Translate(-tileWidth+sideLen/2, tileHeight/2-sideLen/2-sideLen),
				p.Svg.Rotate(30.0, 0.0, sideLen),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 16: // left / center-top square
			transforms = []string{
				p.Svg.Scale(1, -1),
				p.Svg.Translate(sideLen/2, -tileHeight+tileHeight/2-sideLen/2-sideLen),
				p.Svg.Rotate(30.0, 0.0, sideLen),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 17: // right / center-bottom square
			transforms = []string{
				p.Svg.Scale(-1, -1),
				p.Svg.Translate(-tileWidth+sideLen/2, -tileHeight+tileHeight/2-sideLen/2-sideLen),
				p.Svg.Rotate(30.0, 0.0, sideLen),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 18: // left / bottom square
			transforms = []string{
				p.Svg.Scale(1, -1),
				p.Svg.Translate(sideLen/2, -tileHeight+sideLen/2),
				p.Svg.Rotate(-30.0, 0.0, 0.0),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		case 19: // right / bottom square
			transforms = []string{
				p.Svg.Scale(-1, -1),
				p.Svg.Translate(-tileWidth+sideLen/2, -tileHeight+sideLen/2),
				p.Svg.Rotate(-30.0, 0.0, 0.0),
			}
			p.Svg.Rect(0.0, 0.0, sideLen, sideLen, p.tessellationStyles(val, p.Svg.Transform(transforms...))...)
		}
	}
}

func (p *Pattern) tessellationStyles(colourValue float64, style string) []string {
	var styles []string
	styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.Styles.StrokeColour))
	styles = append(styles, fmt.Sprintf(`stroke-opacity="%.2f"`, p.Styles.StrokeOpacity))
	styles = append(styles, fmt.Sprintf(`fill="%s"`, p.fillColour(int(colourValue))))
	styles = append(styles, fmt.Sprintf(`fill-opacity="%f"`, p.opacity(colourValue)))
	styles = append(styles, fmt.Sprintf(`stroke-width="%d"`, 1))

	if style != "" {
		styles = append(styles, style)
	}
	return styles
}
