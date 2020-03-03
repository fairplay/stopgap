# stopgap

stopgap is a static site generator written in Go. It takes Markdown as input, and gives you a static, simple, and responsive website as output. Posts and pages are supported.

stopgap was built upon [simple-website](https://github.com/alexanderte/simple-website) to satisfy my own needs.

## Get it

    go get github.com/fairplay/stopgap

## Initialize website

    mkdir title-of-website
    cd title-of-website
    $GOPATH/bin/stopgap

## Create content

    $EDITOR _sections/start.md
    $EDITOR _sections/end.md
    $EDITOR _sections/congig.md
    $EDITOR _posts/YYYY-MM-DD-title-of-post.md
    $EDITOR _pages/title-of-page.md

## Customize website
    $EDITOR css/styles.css

## Regenerate

    $GOPATH/bin/stopgap

## MIT License

Copyright © 2014–2019 Alexander Teinum, © 2020 Ivan Metalnikov

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
