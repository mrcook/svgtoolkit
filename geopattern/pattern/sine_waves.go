package pattern

import (
	"fmt"
	"math"
)

// SineWaves generator pattern.
func (p Pattern) SineWaves() {
	period := math.Floor(p.reMap(p.seedToInt(0, 1), 0, 15, 100, 400))
	amplitude := math.Floor(p.reMap(p.seedToInt(1, 1), 0, 15, 30, 100))
	waveThickness := math.Floor(p.reMap(p.seedToInt(2, 1), 0, 15, 3, 30))

	// fill the canvas with shapes
	rows := int(float64(p.Height) * 0.05)
	period *= float64(p.Width) / period

	index := 0
	for y := 0; y <= rows; y++ {
		xOffset := period / 4 * 0.7
		path := fmt.Sprintf(
			"M0 %f C %f 0, %f 0, %f %f S %f %f, %f %f S %f 0, %f, %f",
			amplitude,
			xOffset,
			period/2-xOffset,
			period/2,
			amplitude,
			period-xOffset,
			amplitude*2,
			period,
			amplitude,
			period*1.5-xOffset,
			period*1.5,
			amplitude,
		)

		colourValue := p.seedToInt(index, 1)

		var styles []string
		styles = append(styles, fmt.Sprintf(`fill="none"`))
		styles = append(styles, fmt.Sprintf(`stroke="%s"`, p.fillColour(int(colourValue))))
		styles = append(styles, p.Svg.Style(fmt.Sprintf(`opacity: %.3f`, p.opacity(colourValue)), fmt.Sprintf(`stroke-width: %fpx`, waveThickness)))

		group := styles
		group = append(group, p.Svg.Transform(p.Svg.Translate(-period/4, waveThickness*float64(y)-amplitude*1.5)))
		p.Svg.Path(path, group...)

		group = styles
		group = append(group, p.Svg.Transform(p.Svg.Translate(-period/4, waveThickness*float64(y)-amplitude*1.5+waveThickness*36)))
		p.Svg.Path(path, group...)

		index += 1
		if index > 38 {
			index = 0
		}
	}
}
