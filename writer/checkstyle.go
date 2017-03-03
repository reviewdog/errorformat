package writer

import (
	"encoding/xml"
	"fmt"
	"io"
	"sync"

	"github.com/haya14busa/errorformat"
	"github.com/haya14busa/go-checkstyle/checkstyle"
)

// CheckStyle represents checkstyle XML writer. http://checkstyle.sourceforge.net/
type CheckStyle struct {
	mu    sync.Mutex
	files map[string]*checkstyle.File
	w     io.Writer
}

func NewCheckStyle(w io.Writer) *CheckStyle {
	return &CheckStyle{w: w}
}

func (c *CheckStyle) Write(e *errorformat.Entry) error {
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

func (c *CheckStyle) Flush() error {
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
