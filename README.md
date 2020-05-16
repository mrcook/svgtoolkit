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

![book cover example: Optimism by Helen Keller](example.jpg)

_An example image of a book cover created for [epubBooks.com](https://www.epubbooks.com/book/1322-optimism)._


## Toolkit Usage

To generate geo patterns from the command-line:

    $ svgtoolkit pattern -w=400 -h=300 -s="SVG Toolkit"


**As a library**

A new `svg.Canvas` must first be initiated with an `io.Writer` and document
width/height dimensions.

	image := bytes.NewBuffer(nil)
	canvas := svg.New(400, 400, image)

Each document requires the opening and closing `svg` tags in between which
various SVG elements can be generated.

	canvas.DocStart(canvas.ViewBox(0, 0, 400, 400))
	// ...more elements
	canvas.DocClose()

Then output the contents:

	fmt.Println(image.String())

Here is an example using the pattern generator:

    width := 400
    height := 400

	image := bytes.NewBuffer(nil)
	p, _ := pattern.New(width, height, seed, image)

	p.Svg.DocStart(p.Svg.ViewBox(0, 0, width, height))

	if err := p.Generate(); err != nil {
		println(err.Error())
		return
	}

	p.Svg.DocClose()

	fmt.Println(image.String())


### Examples

There are several examples that can be generated with the command-line app:

    $ svgtoolkit examples -name="rect"

See the [example command](cmd/svgtoolkit/examples/examples.go) for the source code of these examples.

There are currently three examples:

* `circle` - 3 overlapping coloured circles.
* `rect`   - 3 overlapping coloured rectangles.
* `use`    - SVG logo, using the `pattern` and `ùse` elements. 


## Installation

    $ go get github.com/mrcook/svgtoolkit
    $ go get github.com/mrcook/svgtoolkit/...


## License

Copyright (c) 2020 Michael R. Cook. All rights reserved.

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
