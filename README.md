## errorformat - Vim 'errorformat' implementation in Go

[![Build Status](https://travis-ci.org/haya14busa/errorformat.svg?branch=master)](https://travis-ci.org/haya14busa/errorformat)
[![Coverage Status](https://coveralls.io/repos/github/haya14busa/errorformat/badge.svg?branch=master)](https://coveralls.io/github/haya14busa/errorformat?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/haya14busa/errorformat)](https://goreportcard.com/report/github.com/haya14busa/errorformat)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/haya14busa/errorformat?status.svg)](https://godoc.org/github.com/haya14busa/errorformat)

errorformat is Vim's quickfix [errorformat](http://vimdoc.sourceforge.net/htmldoc/quickfix.html#errorformat) implementation in golang.

errorformat provides default errorformats for major tools.
You can see defined errorformats [here](https://godoc.org/github.com/haya14busa/errorformat/fmts).
Also, it's easy to [add new errorformat](fmts/README.md) in a similar way to Vim's errorformat.

Note that it's highly compatible with Vim implementation, but it doesn't support Vim regex.

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
Usage: errorformat [flags] [errorformat ...]

errorformat reads compiler/linter/static analyzer result from STDIN, formats
them by given 'errorformat' (90% compatible with Vim's errorformat. :h
errorformat), and outputs formated result to STDOUT.

Example:
        $ echo '/path/to/file:14:28: error message\nfile2:3:4: msg' | errorformat "%f:%l:%c: %m"
        /path/to/file|14 col 28| error message
        file2|3 col 4| msg

        $ golint ./... | errorformat -name=golint

The -f flag specifies an alternate format for the entry, using the
syntax of package template.  The default output is equivalent to -f
'{{.String}}'. The struct being passed to the template is:

        type Entry struct {
                // name of a file
                Filename string
                // line number
                Lnum int
                // column number (first column is 1)
                Col int
                // true: "col" is visual column
                // false: "col" is byte index
                Vcol bool
                // error number
                Nr int
                // search pattern used to locate the error
                Pattern string
                // description of the error
                Text string
                // type of the error, 'E', '1', etc.
                Type rune
                // true: recognized error message
                Valid bool

                // Original error lines (often one line. more than one line for multi-line
                // errorformat. :h errorformat-multi-line)
                Lines []string
        }

Flags:
  -f string
        format template (default "{{.String}}")
  -list
        list defined errorformats
  -name string
        defined errorformat name
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
/path/to/dir/F3.scala|86 col 13 warning| local val in method f is never used
```

### Use cases of 'errorformat' outside Vim
- [haya14busa/reviewdog: A code review dog who keeps your codebase healthy](https://github.com/haya14busa/reviewdog)

## :bird: Author
haya14busa (https://github.com/haya14busa)
