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

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Two images",
		Markdown: `
![Picture of Oscar.](/path/to/cat1.jpg)
![Picture of Luna.](/path/to/cat2.jpg)
Awesome captions about the **kitties**.
`,
		Expected: `<figure>
<img src="/path/to/cat1.jpg" alt="Picture of Oscar.">
<img src="/path/to/cat2.jpg" alt="Picture of Luna.">
<figcaption><p>Awesome captions about the <strong>kitties</strong>.</p></figcaption>
</figure>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Three images",
		Markdown: `
![Picture of Oscar.](/path/to/cat1.jpg)
![Picture of Luna.](/path/to/cat2.jpg)
![Picture of Oreo.](/path/to/cat3.jpg)
Awesome captions about the **kitties**.
`,
		Expected: `<figure>
<img src="/path/to/cat1.jpg" alt="Picture of Oscar.">
<img src="/path/to/cat2.jpg" alt="Picture of Luna.">
<img src="/path/to/cat3.jpg" alt="Picture of Oreo.">
<figcaption><p>Awesome captions about the <strong>kitties</strong>.</p></figcaption>
</figure>
`,
	}, t)

}

func TestFigureWithImageLink(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			figure.Figure.WithImageLink(),
		),
	)
	count := 0

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Simple",
		Markdown: `
![Alt text](https://example.com/image.jpg)
This is caption.
`,
		Expected: `<figure>
<a href="https://example.com/image.jpg">
<img src="https://example.com/image.jpg" alt="Alt text">
</a>
<figcaption><p>This is caption.</p></figcaption>
</figure>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Multi images",
		Markdown: `
![Picture of Oscar.](/path/to/cat1.jpg)
![Picture of Luna.](/path/to/cat2.jpg)
Awesome captions about the **kitties**.
`,
		Expected: `<figure>
<a href="/path/to/cat1.jpg">
<img src="/path/to/cat1.jpg" alt="Picture of Oscar.">
</a>
<a href="/path/to/cat2.jpg">
<img src="/path/to/cat2.jpg" alt="Picture of Luna.">
</a>
<figcaption><p>Awesome captions about the <strong>kitties</strong>.</p></figcaption>
</figure>
`,
	}, t)
}
