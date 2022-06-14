release (ALPHA)
===============

A small package used to parse scene release names

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/release)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/release#license-mit)

## Installation

```bash
go get github.com/peterhellberg/release@latest
```

## Examples

### titles.go

```go
package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/peterhellberg/release"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	titles := []string{}

	for _, f := range files {
		if f.IsDir() {
			r, err := release.Parse(f.Name())
			if err == nil && r.Title != "" {
				titles = append(titles, r.Title)
			}
		}
	}

	sort.Strings(titles)

	for _, title := range titles {
		fmt.Println(title)
	}
}
```

## License (MIT)

*Copyright (C) 2014-2022 [Peter Hellberg](https://c7.se)*

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
