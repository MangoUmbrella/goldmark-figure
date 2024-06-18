// Copyright 2023 The goldmark-figure authors
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package figure_test

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"

	figure "github.com/mangoumbrella/goldmark-figure"
)

func TestFigure(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			figure.Figure,
		),
	)
	count := 0

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Link not image",
		Markdown: `
[Just a link](https://example.com)
This is the paragraph content
`,
		Expected: `
<p><a href="https://example.com">Just a link</a>
This is the paragraph content</p>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Simple",
		Markdown: `
![Alt text](https://example.com/image.jpg)
This is the paragraph content.
This is the continued line.
`,
		Expected: `<figure>
<img src="https://example.com/image.jpg" alt="Alt text">
<figcaption><p>This is the paragraph content.
This is the continued line.</p></figcaption>
</figure>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "With title",
		Markdown: `
![Alt text](https://example.com/image.jpg "Image title")
This is the paragraph content.
This is the continued line.
`,
		Expected: `<figure>
<img src="https://example.com/image.jpg" alt="Alt text" title="Image title">
<figcaption><p>This is the paragraph content.
This is the continued line.</p></figcaption>
</figure>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Image in the middle isn't a figure",
		Markdown: `
Following image is in the middle:
![Alt text](https://example.com/image.jpg)
So this won't be a figure.
`,
		Expected: `
<p>Following image is in the middle:
<img src="https://example.com/image.jpg" alt="Alt text">
So this won't be a figure.</p>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Multiple paragraph content",
		Markdown: `
First paragraph.

![Alt text](https://example.com/image.jpg)
This is the paragraph content.
This is the continued line with **bold**.

Last paragraph.
`,
		Expected: `<p>First paragraph.</p>
<figure>
<img src="https://example.com/image.jpg" alt="Alt text">
<figcaption><p>This is the paragraph content.
This is the continued line with <strong>bold</strong>.</p></figcaption>
</figure>
<p>Last paragraph.</p>
`,
	}, t)

}
