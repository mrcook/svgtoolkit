# SVG Toolkit and GeoPattern generator

A toolkit to aid in SVG document generation, along with built-in geo pattern generators.

The pattern generators found in this library are a Go port of the fantastic
[`geo_pattern`](https://github.com/jasonlong/geo_pattern/) Ruby GEM by Jason Long.
To quote from the GEM:

> Generate beautiful tiling SVG patterns from a string. The string is converted into
> a SHA and a color and pattern are determined based on the values in the hash.

Check out the repository linked above for an example of the generated patterns.

Although SVG Toolkit is intended to be used as a library, there is a simple CLI
command for generating `svg` pattern documents.


## Installation

    go get github.com/mrcook/svgtoolkit

To install the `pattern` command:

    $ go get github.com/mrcook/svgtoolkit/...


## Toolkit Usage

A new `svg.Canvas` must first be initiated with an `io.Writer` and document
width/height dimensions.

	image := bytes.NewBuffer(nil)
	canvas := svg.New(400, 400, image)

Each document requires the opening and closing `svg` tags in between which
various SVG elements can be generated.

	canvas.DocStart(generator.Svg.ViewBox(0, 0, 400, 400))
	// ...more elements
	canvas.DocClose()

Then output the contents:

	fmt.Println(image.String())

An example using the pattern generator might be:

    width := 400
    height := 400

	image := bytes.NewBuffer(nil)
	generator, _ := pattern.New(width, height, seed, image)

	generator.Svg.DocStart(generator.Svg.ViewBox(0, 0, width, height))

	if err := generator.Generate(); err != nil {
		println(err.Error())
		return
	}

	generator.Svg.DocClose()

	fmt.Println(image.String())


## LICENSE

Copyright (c) 2020 Michael R. Cook. All rights reserved.

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
