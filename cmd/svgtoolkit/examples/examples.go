package examples

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/mrcook/svgtoolkit/svg"

	"github.com/google/subcommands"
)

type ExamplesCmd struct {
	name string
}

func (ExamplesCmd) Name() string     { return "examples" }
func (ExamplesCmd) Synopsis() string { return "SVG Toolkit: Examples" }
func (ExamplesCmd) Usage() string {
	return `svgtoolkit examples [-name] <STRING>:
  Print SVG to stdout.

Available examples are:
- circle
- rect
- use

The above examples are taken from dev.w3.org.
`
}

func (e *ExamplesCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&e.name, "name", "", "Example name of SVG image to generate.")
}

func (e ExamplesCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	image := bytes.NewBuffer(nil)

	switch e.name {
	case "circle":
		e.circles(image)
	case "rect":
		e.openClipArt(image)
	case "use":
		e.use(image)
	default:
		fmt.Println("unknown example name")
		fmt.Println()
		fmt.Println(e.Usage())
		return subcommands.ExitUsageError
	}

	fmt.Println(image.String())

	return subcommands.ExitSuccess
}

// https://dev.w3.org/SVG/tools/svgweb/samples/svg-files/circles1.svg
func (e ExamplesCmd) circles(image io.Writer) {
	c := svg.New(500, 400, image)

	c.DocStart(c.ViewBox(0, 0, 500, 400), c.Style(`circle:hover {fill-opacity:0.9;}`))
	c.Title("circles1.svg")

	c.Group(c.Style(`fill-opacity: 0.7`))
	c.Circle("6.5cm", "2cm", "100", c.Transform(c.Translate(0, 50)), c.Style(`fill:red`, `stroke:black`, `stroke-width:0.1cm`))
	c.Circle("6.5cm", "2cm", "100", c.Transform(c.Translate(70, 150)), c.Style(`fill:blue`, `stroke:black`, `stroke-width:0.1cm`))
	c.Circle("6.5cm", "2cm", "100", c.Transform(c.Translate(-70, 150)), c.Style(`fill:green`, `stroke:black`, `stroke-width:0.1cm`))
	c.GroupClose()

	c.DocClose()
}

// https://dev.w3.org/SVG/tools/svgweb/samples/svg-files/open-clipart.svg
func (e ExamplesCmd) openClipArt(image io.Writer) {
	c := svg.New(260, 200, image)
	c.DocStart(c.ViewBox(-80, -20, 260, 200))
	c.Title("open-clipart.svg")

	c.Group(c.Transform(c.Translate(60, 30), c.Rotate(15, 0, 0)))
	c.RectRounded(-4, -4, 126, 126, 10, 10, `fill="#000"`, `fill-opacity="0.2"`)
	c.RectRounded(-13, -13, 126, 126, 10, 10, `fill="#fff"`, `stroke="#ccc"`, `stroke-width="2"`)
	c.RectRounded(0, 0, 100, 100, 10, 10, `fill="#f3e533"`)
	c.GroupClose()

	c.Group()
	c.RectRounded(-4, -4, 126, 126, 10, 10, `fill="#000"`, `fill-opacity="0.2"`)
	c.RectRounded(-13, -13, 126, 126, 10, 10, `fill="#fff"`, `stroke="#ccc"`, `stroke-width="2"`)
	c.RectRounded(0, 0, 100, 100, 10, 10, `fill="#ff7f00"`)
	c.GroupClose()

	c.Group(c.Transform(c.Translate(-60, 60), c.Rotate(-15, 0, 0)))
	c.RectRounded(-4, -4, 126, 126, 10, 10, `fill="#000"`, `fill-opacity="0.2"`)
	c.RectRounded(-13, -13, 126, 126, 10, 10, `fill="#fff"`, `stroke="#ccc"`, `stroke-width="2"`)
	c.RectRounded(0, 0, 100, 100, 10, 10, `fill="#bf0000"`)
	c.GroupClose()

	c.DocClose()
}

// https://dev.w3.org/SVG/tools/svgweb/samples/svg-files/svg.svg
func (e ExamplesCmd) use(image io.Writer) {
	c := svg.New(100, 100, image)

	c.DocStart(c.ViewBox(0, 0, 100, 100))
	c.Title("svg.svg")

	c.Group(`id="gtop"`, `stroke-width="12"`, `stroke="#000"`)
	c.Group(`id="svgstar"`, c.Transform(c.Translate(50, 50)))
	c.Path(`M-27-5a7,7,0,1,0,0,10h54a7,7,0,1,0,0-10z`, `id="svgbar"`)
	c.Use("svgbar", c.Transform(c.Rotate(45, 0, 0)))
	c.Use("svgbar", c.Transform(c.Rotate(90, 0, 0)))
	c.Use("svgbar", c.Transform(c.Rotate(135, 0, 0)))
	c.GroupClose()
	c.GroupClose()
	c.Use("svgstar", `fill="#fb4"`)

	c.DocClose()
}
