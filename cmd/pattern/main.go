package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/mrcook/svgtoolkit"
	"github.com/mrcook/svgtoolkit/geopattern/pattern"
)

func main() {
	var width, height int
	var seed string
	var v bool

	flag.IntVar(&width, "w", 400, "Document width, in pixels, default: 400")
	flag.IntVar(&height, "h", 400, "Document height, in pixels, default: 400")
	flag.StringVar(&seed, "s", "", "String to be used as the pattern/colour seed.")

	flag.BoolVar(&v, "v", false, "Print the current version")

	flag.Usage = func() {
		fmt.Println("SVG Toolkit: Pattern Generator.")
		fmt.Println("Note: patterns will be output to stdout.")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Printf("  %s [OPTIONS]", os.Args[0])
		fmt.Println()
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("  -h, --help    Show this message")
	}
	flag.Parse()

	if v {
		fmt.Println(svgtoolkit.Version)
		os.Exit(0)
	}

	if seed == "" {
		fmt.Println("Error: seed string required.")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

	image := bytes.NewBuffer(nil)
	generator, _ := pattern.New(width, height, seed, image)

	generator.Svg.DocStart(generator.Svg.ViewBox(0, 0, width, height))

	if err := generator.Generate(); err != nil {
		println(err.Error())
		return
	}

	generator.Svg.DocClose()

	fmt.Println(image.String())
}
