# goldmark-figure

[![goldmark-figure Go reference](https://pkg.go.dev/badge/github.com/mangoumbrella/goldmark-figure.svg)](https://pkg.go.dev/github.com/mangoumbrella/goldmark-figure)
[![goldmark-figure test results](https://github.com/mangoumbrella/goldmark-figure/actions/workflows/test.yml/badge.svg?event=push)](https://github.com/mangoumbrella/goldmark-figure/actions/workflows/test.yml/badge.svg?event=push)
[![goldmark-figure Go report card](https://goreportcard.com/badge/github.com/mangoumbrella/goldmark-figure)](https://goreportcard.com/report/github.com/mangoumbrella/goldmark-figure)

[goldmark-figure](https://github.com/MangoUmbrella/goldmark-figure) is a
[goldmark](http://github.com/yuin/goldmark)
extension to parse mardown paragraphs that start with an image into HTML
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
    <figcaption>Awesome caption about <strong>Oscar</strong> the kitty.</figcaption>
</figure>
```

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

See [`figure_test.go`](/figure_test.go) for detailed usages.

# LICENSE

MIT

# Authors

Yilei Yang
