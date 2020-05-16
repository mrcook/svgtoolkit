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
	seed   string
}

func (PatternCmd) Name() string     { return "pattern" }
func (PatternCmd) Synopsis() string { return "SVG Toolkit: Geo Pattern Generator" }
func (PatternCmd) Usage() string {
	return `svgtoolkit pattern [-seed] <STRING>:
  Print SVG to stdout.
`
}

func (p *PatternCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.width, "w", 400, "Width of canvas in pixels (px)")
	f.IntVar(&p.height, "h", 400, "Height of canvas in pixels (px)")
	f.StringVar(&p.seed, "s", "", "Seed string used for generating coloured patterns.")
}

func (p PatternCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.seed == "" {
		fmt.Println("missing seed")
		return subcommands.ExitUsageError
	}
	if p.width < 0 || p.height < 0 {
		fmt.Println("width/height must be positive numbers")
		return subcommands.ExitUsageError
	}

	p.generatePattern()

	return subcommands.ExitSuccess
}

func (p PatternCmd) generatePattern() {
	image := bytes.NewBuffer(nil)
	generator, _ := pattern.New(p.width, p.height, p.seed, image)

	generator.Svg.DocStart(generator.Svg.ViewBox(0, 0, p.width, p.height))

	if err := generator.Generate(); err != nil {
		println(err.Error())
		return
	}

	generator.Svg.DocClose()

	fmt.Println(image.String())
}
