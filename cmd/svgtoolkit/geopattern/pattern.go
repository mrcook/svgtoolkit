package geopattern

import (
	"bytes"
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	"github.com/mrcook/svgtoolkit/geopattern/pattern"
)

type PatternCmd struct {
	width  int
	height int
	zoom   float64
	seed   string
}

func (p *PatternCmd) Name() string     { return "pattern" }
func (p *PatternCmd) Synopsis() string { return "SVG Toolkit: Geo Pattern Generator" }
func (p *PatternCmd) Usage() string {
	return `svgtoolkit pattern [-seed] <STRING>:
  Print SVG to stdout.
`
}

func (p *PatternCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.width, "w", 400, "Width of canvas in pixels (px)")
	f.IntVar(&p.height, "h", 400, "Height of canvas in pixels (px)")
	f.Float64Var(&p.zoom, "z", 1.0, "Zoom scales the pattern size (value must be >= 0.1)")
	f.StringVar(&p.seed, "s", "", "Seed string used for generating coloured patterns.")
}

func (p *PatternCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.seed == "" {
		fmt.Println("missing seed")
		return subcommands.ExitUsageError
	}
	if p.width < 0 || p.height < 0 {
		fmt.Println("width/height must be positive numbers")
		return subcommands.ExitUsageError
	}
	if p.zoom < 0.1 {
		fmt.Println("zoom must be >= 0.1")
		return subcommands.ExitUsageError
	}

	if err := p.generatePattern(); err != nil {
		println(err)
		return subcommands.ExitUsageError
	}

	return subcommands.ExitSuccess
}

func (p *PatternCmd) generatePattern() error {
	image := bytes.NewBuffer(nil)
	generator, _ := pattern.New(p.width, p.height, p.seed, image)
	if err := generator.SetZoom(p.zoom); err != nil {
		return err
	}

	generator.Svg.DocStart(generator.Svg.ViewBox(0, 0, p.width, p.height))

	if err := generator.Generate(); err != nil {
		return err
	}

	generator.Svg.DocClose()

	fmt.Println(image.String())
	return nil
}
