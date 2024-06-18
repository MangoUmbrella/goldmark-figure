// Copyright 2023 The goldmark-figure authors
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package parser

import (
	"regexp"

	fast "github.com/mangoumbrella/goldmark-figure/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var imageRegexp = regexp.MustCompile(`^!\[[^[\]]*\](\([^()]*\)|\[[^[\]]*\])\s*$`)

type figureParagraphTransformer struct {
}

var defaultFigureParagraphTransformer = &figureParagraphTransformer{}

// NewFigureParagraphTransformer returns a new ParagraphTransformer
// that can transform paragraphs into figures.
func NewFigureParagraphTransformer() parser.ParagraphTransformer {
	return defaultFigureParagraphTransformer
}

func (b *figureParagraphTransformer) Transform(node *gast.Paragraph, reader text.Reader, pc parser.Context) {
	lines := node.Lines()
	if lines.Len() < 1 {
		return
	}
	var source = reader.Source()
	var firstSeg = lines.At(0)
	var firstLineStr = firstSeg.Value(source)
	// Here we simply match by regex.
	// But this simple regex ignores image descriptions that contain other links.
	// E.g. ![foo ![bar](/url)](/url2).
	// See CommonMark spec: https://spec.commonmark.org/0.30/#images.
	if !imageRegexp.Match(firstLineStr) {
		return
	}
	figure := fast.NewFigure()
	node.Parent().ReplaceChild(node.Parent(), node, figure)

	figureImage := fast.NewFigureImage()
	figureImage.Lines().Append(lines.At(0))
	figure.AppendChild(figure, figureImage)

	var currentLine = 1
	for currentLine < lines.Len() {
		var currentSeg = lines.At(currentLine)
		var currentLineStr = currentSeg.Value(source)
		if imageRegexp.Match(currentLineStr) {
			// Continued images.
			figureImage := fast.NewFigureImage()
			figureImage.Lines().Append(lines.At(currentLine))
			figure.AppendChild(figure, figureImage)
			currentLine += 1
		} else {
			break
		}
	}

	if currentLine < lines.Len() {
		figureCaption := fast.NewFigureCaption()
		for i := currentLine; i < lines.Len(); i++ {
			seg := lines.At(i)
			if i == lines.Len()-1 {
				// trim last newline(\n)
				seg.Stop = seg.Stop - 1
			}
			figureCaption.Lines().Append(seg)
		}
		figure.AppendChild(figure, figureCaption)
	}
}

type figureASTTransformer struct {
}

var defaultFigureASTTransformer = &figureASTTransformer{}

// NewFigureASTTransformer returns a parser.ASTTransformer for tables.
func NewFigureASTTransformer() parser.ASTTransformer {
	return defaultFigureASTTransformer
}

func (a *figureASTTransformer) Transform(node *gast.Document, reader text.Reader, pc parser.Context) {
}
