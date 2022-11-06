// Package svg provides a toolkit for creating SVG documents.
//
// All SVG documents must be initiated with an io.Writer and width/height
// dimensions. You must also include these elements for the document to be
// valid SVG:
//
//	canvas := svg.New(400, 400, ioWriter)
//	canvas.Declaration()
//	canvas.DocStart()
//	// ...additional elements
//	canvas.DocClose()
//
// Functions for many SVG elements are available, along with a number of
// helper functions for generating the optional element attributes.
//
// It should be noted that many of these function use an interface for their
// position, dimensions, and other core attributes. This provides flexibility
// for the caller at the expense of type safety, allowing for a more fluid
// approach to your SVG creation.
// Here is an example using the rect element:
//
//	canvas.Rect(0, 0, 400, 200)
//	canvas.Rect(0, 0, 101.5, 50.75)
//	canvas.Rect(0, 0, "100%", "50%")
//
// Be aware, the SVG will not render correctly unless valid attributes are used.
// See the SVG specification for details on what values can be used for each element.
package svg

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// Canvas represents a SVG document.
type Canvas struct {
	width  int // Canvas width in px
	height int // Canvas height in px
	writer io.Writer
}

// New returns a SVG instance using the given io.Writer and width/height values.
// The dimensions are expected to be given as `px` values.
func New(w, h int, writer io.Writer) *Canvas {
	return &Canvas{width: w, height: h, writer: writer}
}

// Declaration writes an XML declaration to the canvas.
// This should be the first element in a Canvas.
func (s *Canvas) Declaration() {
	s.println(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>`)
}

// DocStart writes an `<svg>` tag to the canvas.
// It accepts an optional number of string attributes. A common attribute would be viewBox:
//
//	`viewBox="0 0 400 400"`
//
// SVG Reference: http://www.w3.org/TR/SVG11/struct.html#SVGElement
func (s *Canvas) DocStart(attributes ...string) {
	s.Declaration()
	s.print(`<svg`)
	s.printf(` width="%dpx" height="%dpx"`, s.width, s.height)
	if len(attributes) > 0 {
		s.printf(` %s`, strings.Join(attributes, " "))
	}
	s.print(` xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"`)
	s.println(` >`)
}

// DocClose writes the closing `</svg> tag to the canvas.
// This should be the last element in a Canvas.
func (s *Canvas) DocClose() { s.println(`</svg>`) }

// Title of a the document.
//
// SVG Reference: https://www.w3.org/TR/SVG11/struct.html#__svg__SVGDocument__title
func (s *Canvas) Title(str string) { s.Element("title", str) }

// Element is a helper function which writes a XML element to the canvas using the given tag name.
// The the contents will be escaped, and written between a start and end tag.
// It also accepts an optional number of string attributes.
func (s *Canvas) Element(tag, str string, attributes ...string) {
	var attrString string
	if len(attributes) > 0 {
		attrString = " " + strings.Join(attributes, " ")
	}

	s.printf("<%s%s>", tag, attrString)
	xml.Escape(s.writer, []byte(str))
	s.printf("</%s>\n", tag)
}

// VoidElement is a helper function which writes a XML element to the canvas using the given tag name.
// It accepts an optional number of string attributes.
func (s *Canvas) VoidElement(tag string, attributes ...string) {
	if len(attributes) > 0 {
		return
	}

	s.printf("<%s %s />\n", tag, strings.Join(attributes, " "))
}

// Raw is a helper function which writes an arbitrary string to the canvas.
func (s *Canvas) Raw(raw string) {
	_, _ = fmt.Fprintln(s.writer, raw)
}

func (s *Canvas) print(a ...interface{}) {
	_, _ = fmt.Fprint(s.writer, a...)
}

func (s *Canvas) println(a ...interface{}) {
	_, _ = fmt.Fprintln(s.writer, a...)
}

func (s *Canvas) printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(s.writer, format, a...)
}
