// Package writer provides error result writers.
package writer

import (
	"io"

	"github.com/haya14busa/errorformat"
)

// Writer represents error result writer.
type Writer interface {
	Write(*errorformat.Entry) error
}

// BufWriter represents buffered error result writer.
type BufWriter interface {
	Writer
	Flush() error
}

// trackingWriter tracks the last byte written on every write so
// we can avoid printing a newline if one was already written or
// if there is no output at all.
type trackingWriter struct {
	w    io.Writer
	last byte
}

func newTrackingWriter(w io.Writer) *trackingWriter {
	return &trackingWriter{
		w:    w,
		last: '\n',
	}
}

func (t *trackingWriter) Write(p []byte) (n int, err error) {
	n, err = t.w.Write(p)
	if n > 0 {
		t.last = p[n-1]
	}
	return
}

var nl = []byte{'\n'}

// WriteNL writes NL.
func (t *trackingWriter) WriteNL() (int, error) {
	return t.w.Write(nl)
}

// NeedNL returns true if the last byte written is not NL.
func (t *trackingWriter) NeedNL() bool {
	return t.last != '\n'
}
