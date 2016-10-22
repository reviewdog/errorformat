package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/haya14busa/errorformat"
	"github.com/haya14busa/errorformat/fmts"
)

const usageMessage = "" +
	`Usage: errorformat [flags] [errorformat ...]

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
`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	entryFmt = flag.String("f", "{{.String}}", "format template")
	name     = flag.String("name", "", "defined errorformat name")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	errorformats := flag.Args()
	if err := run(os.Stdin, os.Stdout, errorformats, *entryFmt, *name); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, efms []string, entryFmt, name string) error {
	if name != "" {
		f, ok := fmts.DefinedFmts()[name]
		if !ok {
			return fmt.Errorf("%q is not defined", name)
		}
		efms = f.Errorformat
	}

	out := newTrackingWriter(w)

	fm := template.FuncMap{
		"join": strings.Join,
	}
	tmpl, err := template.New("main").Funcs(fm).Parse(entryFmt)
	if err != nil {
		return err
	}

	efm, err := errorformat.NewErrorformat(efms)
	if err != nil {
		return err
	}
	s := efm.NewScanner(r)
	for s.Scan() {
		if err := tmpl.Execute(out, s.Entry()); err != nil {
			return err
		}
		if out.NeedNL() {
			out.WriteNL()
		}
	}
	return nil
}

// TrackingWriter tracks the last byte written on every write so
// we can avoid printing a newline if one was already written or
// if there is no output at all.
type TrackingWriter struct {
	w    io.Writer
	last byte
}

func newTrackingWriter(w io.Writer) *TrackingWriter {
	return &TrackingWriter{
		w:    w,
		last: '\n',
	}
}

func (t *TrackingWriter) Write(p []byte) (n int, err error) {
	n, err = t.w.Write(p)
	if n > 0 {
		t.last = p[n-1]
	}
	return
}

var nl = []byte{'\n'}

// WriteNL writes NL.
func (t *TrackingWriter) WriteNL() (int, error) {
	return t.w.Write(nl)
}

// NeedNL returns true if the last byte written is not NL.
func (t *TrackingWriter) NeedNL() bool {
	return t.last != '\n'
}
