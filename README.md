# goldmark-figure

[![goldmark-figure Go reference](https://pkg.go.dev/badge/github.com/mangoumbrella/goldmark-figure.svg)](https://pkg.go.dev/github.com/mangoumbrella/goldmark-figure)
[![goldmark-figure test results](https://github.com/mangoumbrella/goldmark-figure/actions/workflows/test.yml/badge.svg?event=push)](https://github.com/mangoumbrella/goldmark-figure/actions/workflows/test.yml/badge.svg?event=push)
[![goldmark-figure Go report card](https://goreportcard.com/badge/github.com/mangoumbrella/goldmark-figure)](https://goreportcard.com/report/github.com/mangoumbrella/goldmark-figure)

[goldmark-figure](https://github.com/MangoUmbrella/goldmark-figure) is a
[goldmark](http://github.com/yuin/goldmark)
extension to parse markdown paragraphs that start with an image into HTML
`<figure>` elements. One nice thing is it doesn't use any new markdown
syntaxes.

Example markdown source:

```md
![Picture of Oscar.](/path/to/cat.jpg)
Awesome caption about **Oscar** the kitty.
```

Render result:

```html
<figure>
<img src="/path/to/cat.jpg" alt="Picture of Oscar." />
<figcaption><p>Awesome caption about <strong>Oscar</strong> the kitty.</p></figcaption>
</figure>
```

# Why?

Using dedicated `<figure>` and `<figcaption>` elements makes styling images
with descriptions much easier.
[Here](https://mangobaby.app/parenting-tips/how-to-burp-a-newborn) is an
example:

![Example of an HTML figure with figcaption.](/assets/example.png)

I hear they are also good for SEO.

# Installation

```
go get github.com/mangoumbrella/goldmark-figure
```

# Usage

```go
import (
    "bytes"
    "fmt"

    "github.com/mangoumbrella/goldmark-figure"
    "github.com/yuin/goldmark"
)

func main() {
    markdown := goldmark.New(
        goldmark.WithExtensions(
            figure.Figure,
        ),
    )
    source := `
    ![Picture of Oscar.](/path/to/cat.jpg)
    Awesome caption about **Oscar** the kitty.
    `
    var buf bytes.Buffer
    if err := markdown.Convert([]byte(source), &buf); err != nil {
        panic(err)
    }
    fmt.Print(buf.String())
}
```

## Option to add link to the image

Example:

```go
goldmark.WithExtensions(
    figure.Figure.WithImageLink(),
),
```

Render result:

```html
<figure>
<a href="/path/to/cat.jpg">
<img src="/path/to/cat.jpg" alt="Picture of Oscar." />
</a>
<figcaption>Awesome caption about <strong>Oscar</strong> the kitty.</figcaption>
</figure>
```

See [`figure_test.go`](/figure_test.go) for more examples.

# Changelog

## v1.1.0 (2024-06-18)

* New option to add a link to the image when rendering the figure (see #1).

# LICENSE

MIT

# Authors

Yilei Yang
