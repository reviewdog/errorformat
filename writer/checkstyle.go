package writer

import (
	"encoding/xml"
	"fmt"
	"io"
	"sort"
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
	if e.Nr != 0 && e.Type != 0 {
		checkerr.Source = fmt.Sprintf("%s%d", string(e.Type), e.Nr)
	}
	c.files[e.Filename].Errors = append(c.files[e.Filename].Errors, checkerr)
	return nil
}

func (c *CheckStyle) Flush() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	r := &checkstyle.Result{Version: "1.0"}
	for _, f := range c.files {
		sort.Slice(f.Errors, func(i, j int) bool {
			x, y := f.Errors[i], f.Errors[j]
			if x.Line == y.Line {
				return x.Column < y.Column
			}
			return x.Line < y.Line
		})
		r.Files = append(r.Files, f)
	}
	sort.Slice(r.Files, func(i, j int) bool {
		return r.Files[i].Name < r.Files[j].Name
	})
	fmt.Fprint(c.w, xml.Header)
	e := xml.NewEncoder(c.w)
	e.Indent("", "  ")
	defer c.w.Write(nl)
	return e.Encode(r)
}
