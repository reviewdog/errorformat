## errorformat - Vim 'errorformat' implementation in Go

[![Build Status](https://travis-ci.org/haya14busa/errorformat.svg?branch=master)](https://travis-ci.org/haya14busa/errorformat)
[![Coverage Status](https://coveralls.io/repos/github/haya14busa/errorformat/badge.svg?branch=master)](https://coveralls.io/github/haya14busa/errorformat?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/haya14busa/errorformat)](https://goreportcard.com/report/github.com/haya14busa/errorformat)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/haya14busa/errorformat?status.svg)](https://godoc.org/github.com/haya14busa/errorformat)

errorformat is Vim's quickfix [errorformat](http://vimdoc.sourceforge.net/htmldoc/quickfix.html#errorformat) implementation in golang.

It's highly compatible with Vim implementation, but it doesn't support Vim regex.

### Usage

```go
import "github.com/haya14busa/errorformat"
```

### Example 

#### Code:

```go
in := `
golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
`
efm, _ := errorformat.NewErrorformat([]string{`%f:%l:%c: %m`, `%-G%.%#`})
s := efm.NewScanner(strings.NewReader(in))
for s.Scan() {
    fmt.Println(s.Entry())
}
```

#### Output:

```
golint.new.go|3 col 5| exported var V should have comment or be unexported
golint.new.go|5 col 5| exported var NewError1 should have comment or be unexported
golint.new.go|7 col 1| comment on exported function F should be of the form "F ..."
golint.new.go|11 col 1| comment on exported function F2 should be of the form "F2 ..."
```

### CLI tool

#### Installation

```
go get -u github.com/haya14busa/errorformat/cmd/errorformat
```

#### Usage

```
errorformat [flags] [errorformat ...]
```

```
$ cat testdata/golint.in
golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
$ errorformat -f="file:{{.Filename}} line:{{.Lnum}} col:{{.Col}}" "%f:%l:%c: %m" < testdata/golint.in
file:golint.new.go line:3 col:5
file:golint.new.go line:5 col:5
file:golint.new.go line:7 col:1
file:golint.new.go line:11 col:1
```

```
$ cat testdata/sbt.in
[warn] /path/to/F1.scala:203: local val in method f is never used: (warning smaple 3)
[warn]         val x = 1
[warn]             ^
[warn] /path/to/F1.scala:204: local val in method f is never used: (warning smaple 2)
[warn]   val x = 2
[warn]       ^
[error] /path/to/F2.scala:1093: error: value ++ is not a member of Int
[error]     val x = 1 ++ 2
[error]               ^
[warn] /path/to/dir/F3.scala:83: local val in method f is never used
[warn]         val x = 4
[warn]             ^
[error] /path/to/dir/F3.scala:84: error: value ++ is not a member of Int
[error]         val x = 5 ++ 2
[error]                   ^
[warn] /path/to/dir/F3.scala:86: local val in method f is never used
[warn]         val x = 6
[warn]             ^
$ errorformat "%E[%t%.%+] %f:%l: error: %m" "%A[%t%.%+] %f:%l: %m" "%Z[%.%+] %p^" "%C[%.%+] %.%#" "%-G%.%#" < testdata/sbt.in
/path/to/F1.scala|203 col 13 warning| local val in method f is never used: (warning smaple 3)
/path/to/F1.scala|204 col 7 warning| local val in method f is never used: (warning smaple 2)
/path/to/F2.scala|1093 col 15 error| value &#43;&#43; is not a member of Int
/path/to/dir/F3.scala|83 col 13 warning| local val in method f is never used
/path/to/dir/F3.scala|84 col 19 error| value &#43;&#43; is not a member of Int
/path/to/dir/F3.scala|86 col 26 warning| local val in method f is never used
```

## :bird: Author
haya14busa (https://github.com/haya14busa)
