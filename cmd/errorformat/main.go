package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"text/template"

	"github.com/haya14busa/errorformat"
	"github.com/haya14busa/errorformat/fmts"
	"github.com/haya14busa/go-checkstyle/checkstyle"
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

func main() {
	var (
		entryFmt  = flag.String("f", "{{.String}}", "format template for -w=template")
		writerFmt = flag.String("w", "template", "writer format (template|checkstyle)")
		name      = flag.String("name", "", "defined errorformat name")
		list      = flag.Bool("list", false, "list defined errorformats")
	)
	flag.Usage = usage
	flag.Parse()
	errorformats := flag.Args()
	if err := run(os.Stdin, os.Stdout, errorformats, *writerFmt, *entryFmt, *name, *list); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, efms []string, writerFmt, entryFmt, name string, list bool) error {
	if list {
		fs := fmts.DefinedFmts()
		out := make([]string, 0, len(fs))
		for _, f := range fs {
			out = append(out, fmt.Sprintf("%s\t\t%s - %s", f.Name, f.Description, f.URL))
		}
		sort.Strings(out)
		fmt.Fprintln(w, strings.Join(out, "\n"))
		return nil
	}

	if name != "" {
		f, ok := fmts.DefinedFmts()[name]
		if !ok {
			return fmt.Errorf("%q is not defined", name)
		}
		efms = f.Errorformat
	}

	var writer Writer

	switch writerFmt {
	case "template", "":
		fm := template.FuncMap{
			"join": strings.Join,
		}
		tmpl, err := template.New("main").Funcs(fm).Parse(entryFmt)
		if err != nil {
			return err
		}
		writer = &TemplateWriter{Template: tmpl, Writer: newTrackingWriter(w)}
	case "checkstyle":
		writer = &CheckStyleWriter{w: w}
	default:
		return fmt.Errorf("unknown writer: -w=%v", writerFmt)
	}
	defer func() {
		if err := writer.Flash(); err != nil {
			log.Println(err)
		}
	}()

	efm, err := errorformat.NewErrorformat(efms)
	if err != nil {
		return err
	}
	s := efm.NewScanner(r)
	for s.Scan() {
		if err := writer.Write(s.Entry()); err != nil {
			return err
		}
	}
	return nil
}

type Writer interface {
	Write(*errorformat.Entry) error
	Flash() error
}

type TemplateWriter struct {
	Template *template.Template
	Writer   *TrackingWriter
}

func (t *TemplateWriter) Write(e *errorformat.Entry) error {
	if err := t.Template.Execute(t.Writer, e); err != nil {
		return err
	}
	if t.Writer.NeedNL() {
		if _, err := t.Writer.WriteNL(); err != nil {
			return err
		}
	}
	return nil
}

func (*TemplateWriter) Flash() error {
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

type CheckStyleWriter struct {
	mu    sync.Mutex
	files map[string]*checkstyle.File
	w     io.Writer
}

func (c *CheckStyleWriter) Write(e *errorformat.Entry) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.files == nil {
		c.files = make(map[string]*checkstyle.File)
	}
	if _, ok := c.files[e.Filename]; !ok {
		c.files[e.Filename] = &checkstyle.File{Name: e.Filename}
	}
	checkerr := &checkstyle.Error{
		Column:   e.Col,
		Line:     e.Lnum,
		Message:  e.Text,
		Severity: e.Types(),
	}
	c.files[e.Filename].Errors = append(c.files[e.Filename].Errors, checkerr)
	return nil
}

func (c *CheckStyleWriter) Flash() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	r := &checkstyle.Result{Version: "1.0"}
	for _, f := range c.files {
		r.Files = append(r.Files, f)
	}
	fmt.Fprint(c.w, xml.Header)
	e := xml.NewEncoder(c.w)
	e.Indent("", "  ")
	defer c.w.Write(nl)
	return e.Encode(r)
}
